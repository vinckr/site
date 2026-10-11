// Command draft turns Signal "Note to Self" messages (or a supplied text file)
// into website material: Goodreads links become books, IMDb links become movies,
// recipe links land in recipes.md, and everything else is staged as a feed draft
// for review.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	stagingDir = "_staging"
	notesStore = "_staging/notes.jsonl"
	stateFile  = "_staging/state.json"

	bookData  = "content/_data/book.json"
	movieData = "content/_data/movie.json"
	recipesMD = "content/blog/recipes.md"
	feedMD    = "content/feed.md"
)

// state records which notes have already been turned into drafts. It tracks
// note identity rather than a high-water timestamp, because an imported chat
// export carries historical dates that a timestamp marker would skip.
type state struct {
	Processed []string `json:"processed"`
	LastRun   string   `json:"lastRun"`
}

func loadState() (state, map[string]bool) {
	var s state
	seen := map[string]bool{}
	raw, err := os.ReadFile(stateFile)
	if err != nil {
		return s, seen
	}
	if err := json.Unmarshal(raw, &s); err != nil {
		return state{}, seen
	}
	for _, k := range s.Processed {
		seen[k] = true
	}
	return s, seen
}

func saveState(s state) error {
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile, append(buf, '\n'), 0o644)
}

// noteID is a short stable fingerprint of a note, used to remember that it has
// been processed without storing the whole text again.
func noteID(n Note) string {
	sum := sha256.Sum256([]byte(storeKey(n)))
	return hex.EncodeToString(sum[:8])
}

func main() {
	var (
		source  = flag.String("source", "signal", "where to read notes from: signal or file")
		file    = flag.String("file", "", "path to a text file when -source=file")
		account = flag.String("account", os.Getenv("SIGNAL_ACCOUNT"), "Signal account number, e.g. +49...")
		all     = flag.Bool("all", false, "reprocess every stored note, not just new ones")
		dryRun  = flag.Bool("dry-run", false, "report what would change without writing")
		apply   = flag.String("apply", "", "move the feed block from the given draft file into content/feed.md")
	)
	flag.Parse()

	if *apply != "" {
		if err := applyFeedBlock(*apply, feedMD); err != nil {
			fatal(err)
		}
		fmt.Printf("feed block from %s applied to %s\n", *apply, feedMD)
		return
	}

	if err := os.MkdirAll(stagingDir, 0o755); err != nil {
		fatal(err)
	}

	// 1. Ingest into the durable note store.
	stored, err := loadStore(notesStore)
	if err != nil {
		fatal(err)
	}

	var incoming []Note
	switch *source {
	case "signal":
		incoming, err = drainSignal(*account)
		if err != nil {
			fmt.Fprintf(os.Stderr, "signal: %v\n", err)
			fmt.Fprintf(os.Stderr, "continuing with %d note(s) already in the store\n", len(stored))
		}
	case "file":
		if *file == "" {
			fatal(fmt.Errorf("-source=file requires -file"))
		}
		incoming, err = readFileNotes(*file)
		if err != nil {
			fatal(err)
		}
	default:
		fatal(fmt.Errorf("unknown -source %q (want signal or file)", *source))
	}

	added, err := appendStore(notesStore, stored, incoming)
	if err != nil {
		fatal(err)
	}
	if added > 0 {
		fmt.Printf("captured %d new note(s) into %s\n", added, notesStore)
	}

	notes, err := loadStore(notesStore)
	if err != nil {
		fatal(err)
	}
	sortNotes(notes)

	// 2. Select the notes that have not been processed yet.
	st, processed := loadState()
	var pending []Note
	for _, n := range notes {
		if *all || !processed[noteID(n)] {
			pending = append(pending, n)
		}
	}
	if len(pending) == 0 {
		fmt.Println("nothing new to process")
		return
	}
	fmt.Printf("processing %d note(s)\n", len(pending))

	// 3. Classify and resolve.
	var books, movies, recipeLinks, feed, review, texts, recipeTexts, skipped, ratings []Item
	for _, n := range pending {
		urls, rest := extractURLs(n.Text)

		if rest != "" {
			if looksLikeRecipeText(n.Text) {
				recipeTexts = append(recipeTexts, Item{Kind: KindRecipeText, Body: n.Text, Note: n})
				continue
			}
			// A bare "Title 4/5" is a review, not loose prose. Only treat it as
			// one when the note carries no link of its own.
			if title, score, ok := parseRating(rest); ok && len(urls) == 0 {
				ratings = append(ratings, resolveRating(title, score, n))
				continue
			}
			texts = append(texts, Item{Kind: KindText, Body: rest, Note: n})
		}

		for _, raw := range urls {
			it := classifyURL(raw, n)
			switch it.Kind {
			case KindBook:
				if it.Resolved {
					review = append(review, it)
				} else {
					books = append(books, it)
				}
			case KindMovie:
				if it.Resolved {
					review = append(review, it)
				} else {
					movies = append(movies, it)
				}
			case KindRecipeLink:
				recipeLinks = append(recipeLinks, it)
			case KindSkipped:
				skipped = append(skipped, it)
			default:
				feed = append(feed, it)
			}
		}
	}

	feed = dedupeAgainstFeed(feed)

	// 4. Write.
	date := time.Now().Format("02-01-2006")
	draftPath := filepath.Join(stagingDir, "draft-"+time.Now().Format("2006-01-02")+".md")

	if *dryRun {
		fmt.Println("\n-- dry run, nothing written --")
		report("books", books)
		report("movies", movies)
		report("recipe links", recipeLinks)
		report("feed", feed)
		report("needs review", review)
		report("filtered out", skipped)
		report("movie ratings", ratings)
		fmt.Printf("text blocks: %d, recipe texts: %d\n", len(texts), len(recipeTexts))
		fmt.Printf("draft would be written to %s\n", draftPath)
		return
	}

	existingBooks, err := loadEntries(bookData)
	if err != nil {
		fatal(err)
	}
	nb, err := writeEntries(bookData, existingBooks, toEntries(books))
	if err != nil {
		fatal(err)
	}

	existingMovies, err := loadEntries(movieData)
	if err != nil {
		fatal(err)
	}
	nm, err := writeEntries(movieData, existingMovies, toEntries(movies))
	if err != nil {
		fatal(err)
	}

	nr, err := appendRecipeLinks(recipesMD, recipeLinks)
	if err != nil {
		fatal(err)
	}

	if err := writeDraft(draftPath, date, feed, review, texts, recipeTexts, skipped, ratings); err != nil {
		fatal(err)
	}

	for _, n := range pending {
		if id := noteID(n); !processed[id] {
			processed[id] = true
			st.Processed = append(st.Processed, id)
		}
	}
	st.LastRun = time.Now().Format(time.RFC3339)
	if err := saveState(st); err != nil {
		fatal(err)
	}

	fmt.Printf("\n%s  +%d book(s)\n", bookData, nb)
	fmt.Printf("%s  +%d movie(s)\n", movieData, nm)
	fmt.Printf("%s  +%d recipe link(s)\n", recipesMD, nr)
	fmt.Printf("%s  %d feed link(s), %d to review, %d text block(s), %d recipe text(s), %d filtered out, %d rating(s)\n",
		draftPath, len(feed), len(review), len(texts), len(recipeTexts), len(skipped), len(ratings))
	fmt.Printf("\nreview %s, write your comments, then: make draft-apply\n", draftPath)
}

// classifyURL routes a single URL to its destination, resolving indirect
// sources (a Wikipedia article about a film or book) where it can.
func classifyURL(raw string, n Note) Item {
	// Checked first: these never carry a title worth fetching and the hosts
	// block scrapers anyway, so skip the network round trip entirely.
	if isSkippedHost(raw) {
		return Item{Kind: KindSkipped, URL: raw, Title: fallbackTitle(raw), Note: n}
	}

	if u, ok := normalizeIMDb(raw); ok {
		title, err := imdbTitle(imdbTitleRe.FindStringSubmatch(u)[1])
		if err != nil {
			title = fallbackTitle(raw)
		}
		return Item{Kind: KindMovie, URL: u, Title: title, Note: n}
	}

	if u, ok := normalizeGoodreads(raw); ok {
		title, err := pageTitle(u)
		if err != nil {
			// Goodreads rate-limits aggressively; the slug in the original URL
			// is a better fallback than anything derived from the bare path.
			if slug := goodreadsSlugTitle(raw); slug != "" {
				title = slug
			} else {
				title = fallbackTitle(raw)
			}
		}
		return Item{Kind: KindBook, URL: u, Title: title, Note: n}
	}

	if isWikipedia(raw) {
		if kind, title, target, err := resolveViaWikipedia(raw); err == nil {
			return Item{
				Kind: kind, URL: target, Title: title, Note: n,
				Resolved: true, From: raw,
				Note2: "resolved from Wikipedia via Wikidata",
			}
		}
		// Not a film or book -- fall through to the feed.
	}

	if isRecipeURL(raw) {
		title, err := pageTitle(raw)
		if err != nil {
			title = fallbackTitle(raw)
		}
		return Item{Kind: KindRecipeLink, URL: raw, Title: title, Note: n}
	}

	title, err := pageTitle(raw)
	if err != nil {
		title = fallbackTitle(raw)
	}
	return Item{Kind: KindFeed, URL: raw, Title: title, Note: n}
}

// resolveRating looks the reviewed title up on IMDb so the draft can offer a
// line ready to paste into movies.md. The search can land on the wrong film, so
// a title that does not match what was typed is marked for a second look.
func resolveRating(typed string, score int, n Note) Item {
	it := Item{Kind: KindRating, Title: typed, Rating: score, Note: n, Body: typed}

	found, u, err := imdbSearch(typed)
	if err != nil {
		it.Note2 = "no IMDb match found"
		return it
	}
	it.Title, it.URL = found, u
	if !sameTitle(typed, found) {
		it.Note2 = fmt.Sprintf("matched %q", found)
	}
	return it
}

// sameTitle compares titles loosely, ignoring case, punctuation and spacing.
func sameTitle(a, b string) bool { return titleKey(a) == titleKey(b) }

func titleKey(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// fallbackTitle is used when a page will not give up a title (Cloudflare 403,
// bot challenge), so the item still shows up in the draft with something
// editable rather than being dropped. It reads the last meaningful path
// segment, which is usually the article slug.
func fallbackTitle(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	segs := strings.FieldsFunc(u.Path, func(r rune) bool { return r == '/' })
	for i := len(segs) - 1; i >= 0; i-- {
		seg := segs[i]
		if ext := path.Ext(seg); ext != "" {
			seg = strings.TrimSuffix(seg, ext)
		}
		seg = strings.NewReplacer("-", " ", "_", " ", "+", " ").Replace(seg)
		if unslugged, err := url.PathUnescape(seg); err == nil {
			seg = unslugged
		}
		seg = strings.TrimSpace(seg)
		// Skip bare ids and date fragments; they make useless titles.
		if seg == "" || isNumeric(seg) || len(seg) < 3 {
			continue
		}
		return titleCase(seg)
	}
	if h := host(raw); h != "" {
		return h
	}
	return raw
}

func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return s != ""
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		r := []rune(w)
		if len(r) > 0 && r[0] >= 'a' && r[0] <= 'z' {
			r[0] -= 32
			words[i] = string(r)
		}
	}
	return strings.Join(words, " ")
}

func toEntries(items []Item) []Entry {
	out := make([]Entry, 0, len(items))
	for _, it := range items {
		out = append(out, Entry{Title: it.Title, URL: it.URL})
	}
	return out
}

// dedupeAgainstFeed drops links already present in feed.md and within the batch.
func dedupeAgainstFeed(items []Item) []Item {
	raw, err := os.ReadFile(feedMD)
	existing := ""
	if err == nil {
		existing = string(raw)
	}
	seen := map[string]bool{}
	var out []Item
	for _, it := range items {
		k := strings.ToLower(it.URL)
		if seen[k] || (existing != "" && strings.Contains(existing, it.URL)) {
			continue
		}
		seen[k] = true
		out = append(out, it)
	}
	return out
}

func report(label string, items []Item) {
	fmt.Printf("%s: %d\n", label, len(items))
	for _, it := range items {
		fmt.Printf("  - %s  %s\n", it.Title, it.URL)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
