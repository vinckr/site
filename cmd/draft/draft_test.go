package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestExtractURLsKeepsBalancedParens(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"see https://en.wikipedia.org/wiki/Stalker_(1979_film)", "https://en.wikipedia.org/wiki/Stalker_(1979_film)"},
		{"(https://example.com/a)", "https://example.com/a"},
		{"go to https://example.com/page.", "https://example.com/page"},
		{"https://example.com/x?a=1&b=2, then", "https://example.com/x?a=1&b=2"},
	}
	for _, c := range cases {
		got, _ := extractURLs(c.in)
		if len(got) != 1 || got[0] != c.want {
			t.Errorf("extractURLs(%q) = %v, want [%s]", c.in, got, c.want)
		}
	}
}

func TestExtractURLsLeavesProse(t *testing.T) {
	_, rest := extractURLs("a thought https://example.com and more")
	if !strings.Contains(rest, "a thought") || !strings.Contains(rest, "and more") {
		t.Errorf("prose lost: %q", rest)
	}
}

func TestNormalizeIMDb(t *testing.T) {
	for _, in := range []string{
		"https://m.imdb.com/title/tt0110912/?ref_=nv_sr_1",
		"https://www.imdb.com/title/tt0110912/",
		"imdb.com/title/tt0110912",
		// Language-prefixed share links must land on the same canonical URL.
		"https://www.imdb.com/pt/title/tt0110912/",
		"https://m.imdb.com/de/title/tt0110912/?ref_=ext_shr",
		"https://www.imdb.com/es-es/title/tt0110912/",
	} {
		got, ok := normalizeIMDb(in)
		if !ok || got != "https://www.imdb.com/title/tt0110912/" {
			t.Errorf("normalizeIMDb(%q) = %q, %v", in, got, ok)
		}
	}
	if _, ok := normalizeIMDb("https://www.imdb.com/name/nm0000233/"); ok {
		t.Error("a person page must not be treated as a title")
	}
}

func TestNormalizeGoodreads(t *testing.T) {
	got, ok := normalizeGoodreads("https://www.goodreads.com/book/show/77203.The_Kite_Runner?from_search=true")
	if !ok || got != "https://www.goodreads.com/book/show/77203" {
		t.Errorf("got %q, %v", got, ok)
	}
	if got, ok := normalizeGoodreads("https://www.goodreads.com/list/show/8244"); !ok ||
		got != "https://www.goodreads.com/list/show/8244" {
		t.Errorf("list URL: got %q, %v", got, ok)
	}
	// The language segment is dropped so stored links stay plain goodreads.com.
	for in, want := range map[string]string{
		"https://www.goodreads.com/en/book/show/55428.Goodbye_to_All_That":   "https://www.goodreads.com/book/show/55428",
		"https://www.goodreads.com/en/book/show/50641444-instantiation":      "https://www.goodreads.com/book/show/50641444",
		"https://www.goodreads.com/en/list/show/8244":                        "https://www.goodreads.com/list/show/8244",
		"https://www.goodreads.com/pt-br/book/show/77203.The_Kite_Runner?x=": "https://www.goodreads.com/book/show/77203",
	} {
		if got, ok := normalizeGoodreads(in); !ok || got != want {
			t.Errorf("normalizeGoodreads(%q) = %q, %v; want %q", in, got, ok, want)
		}
	}
}

func TestGoodreadsSlugTitle(t *testing.T) {
	if got := goodreadsSlugTitle("https://www.goodreads.com/book/show/77203.The_Kite_Runner"); got != "The Kite Runner" {
		t.Errorf("got %q", got)
	}
	if got := goodreadsSlugTitle("https://www.goodreads.com/en/book/show/6259.Birdsong"); got != "Birdsong" {
		t.Errorf("language-prefixed URL: got %q", got)
	}
	if got := goodreadsSlugTitle("https://www.goodreads.com/book/show/77203"); got != "" {
		t.Errorf("want empty for slugless URL, got %q", got)
	}
}

func TestInstagramIsFilteredFromFeed(t *testing.T) {
	for _, u := range []string{
		"https://www.instagram.com/p/Cabc123/",
		"https://instagram.com/reel/xyz",
		"https://de.instagram.com/p/Cabc123/",
		"https://instagr.am/p/Cabc123/",
	} {
		if !isSkippedHost(u) {
			t.Errorf("%s was not filtered", u)
		}
		if got := classifyURL(u, Note{}).Kind; got != KindSkipped {
			t.Errorf("classifyURL(%s) = %s, want %s", u, got, KindSkipped)
		}
	}
	// A host that merely contains the word must not be caught.
	for _, u := range []string{
		"https://notinstagram.com/p/x",
		"https://instagram.com.evil.example/p/x",
		"https://simonwillison.net/2025/Jan/1/",
	} {
		if isSkippedHost(u) {
			t.Errorf("%s was filtered but should not be", u)
		}
	}
}

func TestSkippedItemsReachDraftNotFeed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "draft.md")
	feed := []Item{{Kind: KindFeed, URL: "https://example.com/good", Title: "Good"}}
	skipped := []Item{{Kind: KindSkipped, URL: "https://www.instagram.com/p/C1/", Title: "P"}}
	if err := writeDraft(path, "01-10-2026", feed, nil, nil, nil, skipped, nil); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(path)
	s := string(out)
	if !strings.Contains(s, "## Filtered out") || !strings.Contains(s, "instagram.com/p/C1") {
		t.Error("skipped link missing from draft")
	}
	// It must not sit inside the block that gets applied to feed.md.
	start := strings.Index(s, "<!-- feed:begin -->")
	end := strings.Index(s, "<!-- feed:end -->")
	if start == -1 || end == -1 {
		t.Fatal("feed block markers missing")
	}
	block := s[start:end]
	if strings.Contains(block, "instagram") {
		t.Error("skipped link leaked into the feed block")
	}
	if !strings.Contains(block, "example.com/good") {
		t.Error("ordinary link missing from the feed block")
	}
}

func TestRecipeDetection(t *testing.T) {
	if !isRecipeURL("https://www.seriouseats.com/the-best-chili") {
		t.Error("known recipe host not detected")
	}
	if !isRecipeURL("https://someblog.com/recipes/apple-cake") {
		t.Error("recipe path not detected")
	}
	if isRecipeURL("https://simonwillison.net/2025/Jan/1/") {
		t.Error("ordinary link misrouted to recipes")
	}

	recipe := "Apple Cake\nIngredients\n- apples\nSteps\n1. bake"
	if !looksLikeRecipeText(recipe) {
		t.Error("recipe text not detected")
	}
	if looksLikeRecipeText("I had a great cake today, the ingredients were a mystery") {
		t.Error("prose mentioning ingredients misrouted to recipes")
	}
}

// TestGuardPathRefusesDrafts is the important one: the _drafts folder is off
// limits to this tool under any circumstances.
func TestGuardPathRefusesDrafts(t *testing.T) {
	for _, p := range []string{
		"_drafts/notes.md",
		"./_drafts/deep/nested.md",
		"content/../_drafts/x.md",
	} {
		if err := guardPath(p); err == nil {
			t.Errorf("guardPath(%q) allowed a write into _drafts", p)
		}
	}
	if err := guardPath("content/feed.md"); err != nil {
		t.Errorf("guardPath rejected a legitimate path: %v", err)
	}
}

func TestWriteEntriesDedupesAndPreservesAmpersand(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "book.json")

	existing := []Entry{{Title: "Sex, Drugs & Magick", URL: "https://www.goodreads.com/book/show/1"}}
	added := []Entry{
		{Title: "Dupe", URL: "https://www.goodreads.com/book/show/1"}, // same URL, must be dropped
		{Title: "A New Book", URL: "https://www.goodreads.com/book/show/2"},
	}
	n, err := writeEntries(path, existing, added)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Errorf("added %d entries, want 1", n)
	}
	out, _ := os.ReadFile(path)
	escaped := "\\" + "u0026" // the literal six characters \u0026
	if strings.Contains(string(out), escaped) {
		t.Errorf("ampersand was HTML-escaped: %s", out)
	}
	if !strings.Contains(string(out), "Sex, Drugs & Magick") {
		t.Error("existing title mangled")
	}
	// sorted by title: "A New Book" must precede "Sex, Drugs & Magick"
	if strings.Index(string(out), "A New Book") > strings.Index(string(out), "Sex, Drugs") {
		t.Error("entries not sorted by title")
	}
}

func TestApplyFeedBlockInsertsAboveNewestEntry(t *testing.T) {
	dir := t.TempDir()
	draft := filepath.Join(dir, "draft.md")
	feed := filepath.Join(dir, "feed.md")

	os.WriteFile(draft, []byte("x\n<!-- feed:begin -->\n### 30-09-2026\n\n- [A](https://a.example) - nice\n<!-- feed:end -->\ny\n"), 0o644)
	os.WriteFile(feed, []byte("---\ntitle: feed\n---\n\nintro\n\n### 10-03-2026\n\n- [Old](https://old.example)\n"), 0o644)

	if err := applyFeedBlock(draft, feed); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(feed)
	s := string(got)
	if !strings.Contains(s, "### 30-09-2026") {
		t.Fatal("new block missing")
	}
	if strings.Index(s, "### 30-09-2026") > strings.Index(s, "### 10-03-2026") {
		t.Error("new entry was not placed above the previous newest entry")
	}
	if !strings.Contains(s, "intro") || !strings.Contains(s, "title: feed") {
		t.Error("front matter or intro was damaged")
	}
}

func TestDedupeStoreKeyBySource(t *testing.T) {
	a := Note{Timestamp: 1, Text: "hello", Source: "file"}
	b := Note{Timestamp: 999, Text: "hello", Source: "file"}
	if storeKey(a) != storeKey(b) {
		t.Error("file notes with the same text must dedupe regardless of timestamp")
	}
	c := Note{Timestamp: 1, Text: "hello", Source: "signal"}
	d := Note{Timestamp: 999, Text: "hello", Source: "signal"}
	if storeKey(c) == storeKey(d) {
		t.Error("signal notes on different days must stay distinct")
	}
}

func TestFallbackTitleFromSlug(t *testing.T) {
	if got := fallbackTitle("https://www.seriouseats.com/the-best-chili-recipe"); got != "The Best Chili Recipe" {
		t.Errorf("got %q", got)
	}
}

func TestParseDateHeader(t *testing.T) {
	ref := time.Date(2026, time.October, 8, 12, 0, 0, 0, time.Local)
	cases := []struct {
		in   string
		want string // yyyy-mm-dd, empty means "not a date header"
	}{
		{"Fri 10 Apr", "2026-04-10"},
		{"Sat 11 Apr", "2026-04-11"},
		{"13 Mar 2026", "2026-03-13"},
		{"Mon, 5 May 2025", "2025-05-05"},
		{"10 April", "2026-04-10"},
		{"Mar 13, 2026", "2026-03-13"},
		{"Today", "2026-10-08"},
		{"Yesterday", "2026-10-07"},
		{"13/03/2026", "2026-03-13"},
		// A December header with no year, read in October, is last December.
		{"25 Dec", "2025-12-25"},
		// Not date headers:
		{"https://example.com", ""},
		{"14:37", ""},
		{"Horrible Bosses 3/5", ""},
		{"31 Feb", ""},
		{"The studi", ""},
	}
	for _, c := range cases {
		got, ok := parseDateHeader(c.in, ref)
		if c.want == "" {
			if ok {
				t.Errorf("parseDateHeader(%q) was accepted as %s", c.in, got.Format("2006-01-02"))
			}
			continue
		}
		if !ok {
			t.Errorf("parseDateHeader(%q) not recognised", c.in)
			continue
		}
		if got.Format("2006-01-02") != c.want {
			t.Errorf("parseDateHeader(%q) = %s, want %s", c.in, got.Format("2006-01-02"), c.want)
		}
	}
}

func TestParseTimeOfDay(t *testing.T) {
	for _, c := range []struct {
		in   string
		want time.Duration
		ok   bool
	}{
		{"14:37", 14*time.Hour + 37*time.Minute, true},
		{"3:30", 3*time.Hour + 30*time.Minute, true},
		{"9:05 PM", 21*time.Hour + 5*time.Minute, true},
		{"12:00 am", 0, true},
		{"25:00", 0, false},
		{"3/5", 0, false},
		{"hello", 0, false},
	} {
		got, ok := parseTimeOfDay(c.in)
		if ok != c.ok || (ok && got != c.want) {
			t.Errorf("parseTimeOfDay(%q) = %v, %v; want %v, %v", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestReadFileNotesSplitsByDayAndMessage(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "export.txt")
	os.WriteFile(path, []byte(
		"Fri 10 Apr\n\nhttps://a.example/\n14:37\n\nhttp://b.example/\n14:44\n\n"+
			"Sat 11 Apr\n\nhttps://c.example/\n17:40\n"), 0o644)

	notes, err := readFileNotes(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) != 3 {
		t.Fatalf("got %d notes, want 3: %+v", len(notes), notes)
	}
	// Times must be consumed, not left in the text.
	for _, n := range notes {
		if strings.Contains(n.Text, ":") && !strings.Contains(n.Text, "//") {
			t.Errorf("time line leaked into note text: %q", n.Text)
		}
		if strings.Contains(n.Text, "\n") {
			t.Errorf("note should be a single line: %q", n.Text)
		}
	}
	if d := notes[0].When().Format("2006-01-02"); d != "2026-04-10" && d != "2025-04-10" {
		t.Errorf("first note dated %s, want 10 April", d)
	}
	if notes[0].When().Format("15:04") != "14:37" {
		t.Errorf("time of day = %s, want 14:37", notes[0].When().Format("15:04"))
	}
	if notes[2].When().Day() != 11 {
		t.Errorf("third note should fall on the 11th, got %s", notes[2].When())
	}
}

func TestRenderFeedDaysGroupsNewestFirst(t *testing.T) {
	day := func(d int) Note {
		return Note{Timestamp: time.Date(2026, time.April, d, 12, 0, 0, 0, time.Local).UnixMilli()}
	}
	feed := []Item{
		{Title: "A", URL: "https://a.example", Note: day(10)},
		{Title: "B", URL: "https://b.example", Note: day(10)},
		{Title: "C", URL: "https://c.example", Note: day(11)},
	}
	got := renderFeedDays(feed)

	if strings.Count(got, "### ") != 2 {
		t.Errorf("want 2 day headings, got:\n%s", got)
	}
	if !strings.Contains(got, "### 10-04-2026") || !strings.Contains(got, "### 11-04-2026") {
		t.Errorf("headings wrong:\n%s", got)
	}
	if strings.Index(got, "### 11-04-2026") > strings.Index(got, "### 10-04-2026") {
		t.Error("newest day must come first")
	}
	// A and B belong to the 10th, under its heading and after the 11th's block.
	tenth := got[strings.Index(got, "### 10-04-2026"):]
	if !strings.Contains(tenth, "A") || !strings.Contains(tenth, "B") {
		t.Errorf("links not grouped under their day:\n%s", got)
	}
	if strings.Contains(tenth, "C") {
		t.Error("link from the 11th leaked into the 10th")
	}
}

func TestApplyFeedBlockCarriesEveryDay(t *testing.T) {
	dir := t.TempDir()
	draft := filepath.Join(dir, "draft.md")
	feed := filepath.Join(dir, "feed.md")
	os.WriteFile(draft, []byte("<!-- feed:begin -->\n### 11-04-2026\n\n- [C](https://c.example) - x\n\n### 10-04-2026\n\n- [A](https://a.example) - y\n<!-- feed:end -->\n"), 0o644)
	os.WriteFile(feed, []byte("---\ntitle: feed\n---\n\n### 01-01-2026\n\n- [Old](https://old.example)\n"), 0o644)

	if err := applyFeedBlock(draft, feed); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(feed)
	s := string(got)
	for _, want := range []string{"### 11-04-2026", "### 10-04-2026", "### 01-01-2026"} {
		if !strings.Contains(s, want) {
			t.Errorf("missing %s after apply:\n%s", want, s)
		}
	}
	if strings.Index(s, "### 11-04-2026") > strings.Index(s, "### 01-01-2026") {
		t.Error("new days must be inserted above the existing newest entry")
	}
}

func TestParseRating(t *testing.T) {
	cases := []struct {
		in    string
		title string
		score int
		ok    bool
	}{
		{"Horrible Bosses 3/5", "Horrible Bosses", 3, true},
		{"Bugonia 4/5", "Bugonia", 4, true},
		{"Synecdoche, New York 5/5", "Synecdoche, New York", 5, true},
		{"The Whistler 3 / 5", "The Whistler", 3, true},
		{"Dune 5/5.", "Dune", 5, true},
		// Not ratings:
		{"3/5", "", 0, false},   // no title
		{"A 3/5", "", 0, false}, // title too short to be real
		{"Some film 6/5", "", 0, false},
		{"Some film 0/5", "", 0, false},
		{"meeting at 3/5 pm", "", 0, false},
		{"Apple Cake\nIngredients 3/5", "", 0, false}, // multi-line
		{"just a thought", "", 0, false},
	}
	for _, c := range cases {
		title, score, ok := parseRating(c.in)
		if ok != c.ok || title != c.title || score != c.score {
			t.Errorf("parseRating(%q) = %q, %d, %v; want %q, %d, %v",
				c.in, title, score, ok, c.title, c.score, c.ok)
		}
	}
}

func TestBandcampIsFiltered(t *testing.T) {
	for _, u := range []string{
		"https://polygonia.bandcamp.com/album/ceaseless-motion",
		"https://bandcamp.com/something",
	} {
		if !isSkippedHost(u) {
			t.Errorf("%s was not filtered", u)
		}
	}
	// YouTube stays in the feed: it is decided case by case.
	for _, u := range []string{
		"https://www.youtube.com/watch?v=abc",
		"https://youtu.be/abc",
		"https://m.youtube.com/watch?v=abc",
	} {
		if isSkippedHost(u) {
			t.Errorf("%s must not be filtered", u)
		}
	}
}

func TestRatingsSectionIsPasteReady(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "draft.md")
	ratings := []Item{
		{Kind: KindRating, Title: "Horrible Bosses", URL: "https://www.imdb.com/title/tt1499658/", Rating: 3, Body: "Horrible Bosses"},
		{Kind: KindRating, Title: "Unknown Thing", Rating: 2, Body: "Unknown Thing", Note2: "no IMDb match found"},
	}
	if err := writeDraft(path, "11-10-2026", nil, nil, nil, nil, nil, ratings); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(path)
	s := string(out)

	want := "- [Horrible Bosses](https://www.imdb.com/title/tt1499658/) 3/5"
	if !strings.Contains(s, want) {
		t.Errorf("missing movies.md-formatted line %q in:\n%s", want, s)
	}
	// An unresolved title still appears, just without a link.
	if !strings.Contains(s, "- Unknown Thing 2/5") {
		t.Errorf("unresolved rating dropped:\n%s", s)
	}
	if !strings.Contains(s, "no IMDb match found") {
		t.Error("unresolved rating was not flagged for a second look")
	}
	// Ratings must never reach the feed block.
	if strings.Contains(s, "<!-- feed:begin -->") {
		t.Error("unexpected feed block")
	}
}

func TestJSONEntryFragmentMatchesDataFileStyle(t *testing.T) {
	got := jsonEntryFragment(Entry{
		Title: "The Cook, the Thief, His Wife & Her Lover",
		URL:   "https://www.imdb.com/title/tt0097108/",
	})
	want := "  {\n" +
		"    \"title\": \"The Cook, the Thief, His Wife & Her Lover\",\n" +
		"    \"url\": \"https://www.imdb.com/title/tt0097108/\"\n" +
		"  },\n"
	if got != want {
		t.Errorf("fragment does not match data file style.\ngot:\n%s\nwant:\n%s", got, want)
	}
	escaped := "\\" + "u0026"
	if strings.Contains(got, escaped) {
		t.Error("ampersand was escaped")
	}
}

// TestReviewFragmentsArePasteable proves the rendered block really can be
// dropped into the data file: splicing it into an existing array must parse.
func TestReviewFragmentsArePasteable(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "draft.md")
	review := []Item{
		{Kind: KindMovie, Title: "Stalker", URL: "https://www.imdb.com/title/tt0079944/",
			From: "https://en.wikipedia.org/wiki/Stalker_(1979_film)", Resolved: true,
			Note2: "resolved from Wikipedia via Wikidata"},
		{Kind: KindBook, Title: "Dune", URL: "https://www.goodreads.com/book/show/44767458",
			From: "https://en.wikipedia.org/wiki/Dune_(novel)", Resolved: true},
	}
	if err := writeDraft(path, "11-10-2026", nil, review, nil, nil, nil, nil); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(path)
	s := string(out)

	// Each destination gets its own block.
	if !strings.Contains(s, "### "+movieData) || !strings.Contains(s, "### "+bookData) {
		t.Fatalf("missing per-file sections:\n%s", s)
	}
	// The movie entry must not appear under the book heading.
	bookPart := s[strings.Index(s, "### "+bookData):]
	if strings.Contains(bookPart, "imdb.com") {
		t.Error("a movie leaked into the book block")
	}

	// Splice the first fenced block into an array and parse it.
	start := strings.Index(s, "```json\n") + len("```json\n")
	end := start + strings.Index(s[start:], "```")
	fragment := s[start:end]

	var entries []Entry
	spliced := "[\n" + fragment + "  {\"title\": \"Existing\", \"url\": \"https://x.example\"}\n]"
	if err := json.Unmarshal([]byte(spliced), &entries); err != nil {
		t.Fatalf("pasted fragment does not parse: %v\n%s", err, spliced)
	}
	if len(entries) != 2 || entries[0].Title != "Stalker" {
		t.Errorf("unexpected entries after paste: %+v", entries)
	}

	// Provenance stays outside the fenced block so the paste is clean.
	if strings.Contains(fragment, "wikipedia.org") {
		t.Error("provenance leaked into the pasteable JSON")
	}
	if !strings.Contains(s, "from https://en.wikipedia.org/wiki/Stalker_(1979_film)") {
		t.Error("provenance missing from the draft")
	}
}
