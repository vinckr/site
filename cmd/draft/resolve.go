package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 " +
	"(KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"

var client = &http.Client{Timeout: 25 * time.Second}

// throttle keeps requests to the same host spaced out. Goodreads in particular
// serves one page then answers 202 with an empty body for a while, so we pace
// ourselves rather than hammering and falling back to a worse title.
var (
	throttleMu   sync.Mutex
	lastHostHit  = map[string]time.Time{}
	hostInterval = map[string]time.Duration{
		"goodreads.com": 3 * time.Second,
		"imdb.com":      1 * time.Second,
	}
)

func waitForHost(raw string) {
	h := host(raw)
	gap, ok := hostInterval[h]
	if !ok {
		gap = 300 * time.Millisecond
	}
	throttleMu.Lock()
	defer throttleMu.Unlock()
	if last, seen := lastHostHit[h]; seen {
		if wait := gap - time.Since(last); wait > 0 {
			time.Sleep(wait)
		}
	}
	lastHostHit[h] = time.Now()
}

// fetch retrieves a page, retrying the soft bot-challenge responses that IMDb
// and Goodreads hand out intermittently (202 with an empty body, 429, 503).
func fetch(raw string) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt*attempt) * 2 * time.Second)
		}
		waitForHost(raw)
		body, retryable, err := fetchOnce(raw)
		if err == nil {
			return body, nil
		}
		lastErr = err
		if !retryable {
			return nil, err
		}
	}
	return nil, lastErr
}

func fetchOnce(raw string) (body []byte, retryable bool, err error) {
	req, err := http.NewRequest(http.MethodGet, raw, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/json;q=0.9,*/*;q=0.8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, true, err
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == http.StatusOK:
	case resp.StatusCode == http.StatusAccepted,
		resp.StatusCode == http.StatusTooManyRequests,
		resp.StatusCode == http.StatusServiceUnavailable:
		return nil, true, fmt.Errorf("%s: HTTP %d (bot challenge)", raw, resp.StatusCode)
	default:
		return nil, false, fmt.Errorf("%s: HTTP %d", raw, resp.StatusCode)
	}

	b, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, true, err
	}
	if len(b) == 0 {
		return nil, true, fmt.Errorf("%s: empty body", raw)
	}
	return b, false, nil
}

var ogTitleRe = regexp.MustCompile(`<meta[^>]+property=["']og:title["'][^>]+content=["']([^"']+)["']`)

// pageTitle returns a page's og:title, falling back to <title>.
func pageTitle(raw string) (string, error) {
	body, err := fetch(raw)
	if err != nil {
		return "", err
	}
	if m := ogTitleRe.FindSubmatch(body); m != nil {
		return cleanTitle(string(m[1])), nil
	}
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	if t, ok := findTitle(doc); ok {
		return cleanTitle(t), nil
	}
	return "", fmt.Errorf("no title found for %s", raw)
}

func findTitle(n *html.Node) (string, bool) {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data, true
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if t, ok := findTitle(c); ok {
			return t, ok
		}
	}
	return "", false
}

var siteSuffixRe = regexp.MustCompile(`\s*[|\-–—]\s*(Goodreads|IMDb|Wikipedia)\s*$`)

func cleanTitle(s string) string {
	s = html.UnescapeString(strings.TrimSpace(s))
	s = siteSuffixRe.ReplaceAllString(s, "")
	return strings.TrimSpace(strings.Join(strings.Fields(s), " "))
}

// --- IMDb -------------------------------------------------------------------

type imdbSuggestion struct {
	D []struct {
		L  string `json:"l"` // label / title
		ID string `json:"id"`
		Q  string `json:"q"` // qualifier, e.g. "TV series"
		Y  int    `json:"y"` // year
	} `json:"d"`
}

// imdbTitle looks a title up by its tt id. IMDb blocks plain HTTP clients on the
// main site (HTTP 202, empty body), so we use the public suggestion endpoint the
// site's own search box calls.
func imdbTitle(ttID string) (string, error) {
	u := fmt.Sprintf("https://v2.sg.media-imdb.com/suggestion/t/%s.json", ttID)
	body, err := fetch(u)
	if err != nil {
		return "", err
	}
	var s imdbSuggestion
	if err := json.Unmarshal(body, &s); err != nil {
		return "", err
	}
	for _, d := range s.D {
		if d.ID == ttID {
			return imdbLabel(d.L, d.Q), nil
		}
	}
	if len(s.D) > 0 {
		return imdbLabel(s.D[0].L, s.D[0].Q), nil
	}
	return "", fmt.Errorf("no IMDb entry for %s", ttID)
}

// imdbLabel matches the existing movie.json convention of marking series.
func imdbLabel(label, qualifier string) string {
	label = cleanTitle(label)
	if strings.Contains(strings.ToLower(qualifier), "series") {
		label += " (Series)"
	}
	return label
}

// imdbSearch finds the best IMDb match for a free-text title.
func imdbSearch(query string) (title, ttURL string, err error) {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return "", "", fmt.Errorf("empty query")
	}
	u := fmt.Sprintf("https://v3.sg.media-imdb.com/suggestion/x/%s.json?includeVideos=0",
		url.PathEscape(q))
	body, err := fetch(u)
	if err != nil {
		return "", "", err
	}
	var s imdbSuggestion
	if err := json.Unmarshal(body, &s); err != nil {
		return "", "", err
	}
	for _, d := range s.D {
		if strings.HasPrefix(d.ID, "tt") {
			return imdbLabel(d.L, d.Q), "https://www.imdb.com/title/" + d.ID + "/", nil
		}
	}
	return "", "", fmt.Errorf("no IMDb match for %q", query)
}

// --- Goodreads --------------------------------------------------------------

var grSearchRe = regexp.MustCompile(`/book/show/(\d+)`)

// goodreadsSearch finds the best Goodreads book match for a free-text title.
// The first search hit is not always the right edition, so callers stage the
// result for review rather than writing it straight into book.json.
func goodreadsSearch(query string) (title, bookURL string, err error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return "", "", fmt.Errorf("empty query")
	}
	u := "https://www.goodreads.com/search?q=" + url.QueryEscape(q)
	body, err := fetch(u)
	if err != nil {
		return "", "", err
	}
	m := grSearchRe.FindSubmatch(body)
	if m == nil {
		return "", "", fmt.Errorf("no Goodreads match for %q", query)
	}
	bookURL = "https://www.goodreads.com/book/show/" + string(m[1])
	title, err = pageTitle(bookURL)
	if err != nil {
		return "", bookURL, err
	}
	return title, bookURL, nil
}

// --- Wikipedia / Wikidata ---------------------------------------------------

// wikidataItem resolves a Wikipedia article URL to its Wikidata entity id.
func wikidataItem(articleURL string) (string, string, error) {
	u, err := url.Parse(articleURL)
	if err != nil {
		return "", "", err
	}
	title := strings.TrimPrefix(u.Path, "/wiki/")
	if title == "" || title == u.Path {
		return "", "", fmt.Errorf("not a Wikipedia article URL: %s", articleURL)
	}
	title, err = url.PathUnescape(title)
	if err != nil {
		return "", "", err
	}
	lang := strings.SplitN(u.Hostname(), ".", 2)[0]
	api := fmt.Sprintf(
		"https://%s.wikipedia.org/w/api.php?action=query&redirects=1&prop=pageprops&titles=%s&format=json",
		lang, url.QueryEscape(strings.ReplaceAll(title, "_", " ")))

	body, err := fetch(api)
	if err != nil {
		return "", "", err
	}
	var r struct {
		Query struct {
			Pages map[string]struct {
				Title     string `json:"title"`
				PageProps struct {
					WikibaseItem string `json:"wikibase_item"`
				} `json:"pageprops"`
			} `json:"pages"`
		} `json:"query"`
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return "", "", err
	}
	for _, p := range r.Query.Pages {
		if p.PageProps.WikibaseItem != "" {
			return p.PageProps.WikibaseItem, p.Title, nil
		}
	}
	return "", "", fmt.Errorf("no Wikidata item for %s", articleURL)
}

// wikidataClaim returns the first string value of a property on an entity.
func wikidataClaim(entity, property string) (string, error) {
	api := fmt.Sprintf(
		"https://www.wikidata.org/w/api.php?action=wbgetclaims&entity=%s&property=%s&format=json",
		url.QueryEscape(entity), url.QueryEscape(property))
	body, err := fetch(api)
	if err != nil {
		return "", err
	}
	var r struct {
		Claims map[string][]struct {
			Mainsnak struct {
				DataValue struct {
					Value json.RawMessage `json:"value"`
				} `json:"datavalue"`
			} `json:"mainsnak"`
		} `json:"claims"`
	}
	if err := json.Unmarshal(body, &r); err != nil {
		return "", err
	}
	for _, claims := range r.Claims {
		for _, c := range claims {
			var s string
			if err := json.Unmarshal(c.Mainsnak.DataValue.Value, &s); err == nil && s != "" {
				return s, nil
			}
		}
	}
	return "", fmt.Errorf("no %s claim on %s", property, entity)
}

const (
	propIMDbID        = "P345"
	propGoodreadsWork = "P8383"
	propGoodreadsBook = "P2969"
	propInstanceOf    = "P31"
)

// resolveViaWikipedia tries to turn a Wikipedia article about a film or book
// into the corresponding IMDb or Goodreads link.
func resolveViaWikipedia(articleURL string) (kind Kind, title, target string, err error) {
	entity, pageTitleText, err := wikidataItem(articleURL)
	if err != nil {
		return "", "", "", err
	}

	// A film carries an IMDb id directly -- this is exact, not a search guess.
	if ttID, err := wikidataClaim(entity, propIMDbID); err == nil && strings.HasPrefix(ttID, "tt") {
		t, terr := imdbTitle(ttID)
		if terr != nil {
			t = pageTitleText
		}
		return KindMovie, t, "https://www.imdb.com/title/" + ttID + "/", nil
	}

	// A book may carry a Goodreads edition id, which maps straight to /book/show/.
	if bookID, err := wikidataClaim(entity, propGoodreadsBook); err == nil && bookID != "" {
		target := "https://www.goodreads.com/book/show/" + bookID
		t, terr := pageTitle(target)
		if terr != nil {
			t = pageTitleText
		}
		return KindBook, t, target, nil
	}

	// A Goodreads *work* id does not map to a book/show id, so fall back to
	// searching Goodreads by the article title.
	if _, err := wikidataClaim(entity, propGoodreadsWork); err == nil {
		t, u, serr := goodreadsSearch(pageTitleText)
		if serr == nil {
			return KindBook, t, u, nil
		}
	}

	return "", "", "", fmt.Errorf("no film or book identifiers on %s", articleURL)
}

var grSlugRe = regexp.MustCompile(`goodreads\.com/book/show/\d+[.\-]([^/?#]+)`)

// goodreadsSlugTitle recovers a readable title from the URL slug, e.g.
// ".../book/show/77203.The_Kite_Runner" -> "The Kite Runner". Goodreads
// rate-limits hard, so this is often the only title available.
func goodreadsSlugTitle(raw string) string {
	m := grSlugRe.FindStringSubmatch(raw)
	if m == nil {
		return ""
	}
	slug := strings.NewReplacer("_", " ", "-", " ").Replace(m[1])
	return strings.TrimSpace(strings.Join(strings.Fields(slug), " "))
}
