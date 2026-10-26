package main

import "time"

// Note is one raw message captured from Signal Note-to-Self or a supplied text file.
type Note struct {
	Timestamp int64  `json:"timestamp"` // unix milliseconds
	Text      string `json:"text"`
	Source    string `json:"source"` // "signal" or "file"
}

func (n Note) When() time.Time { return time.UnixMilli(n.Timestamp) }

// Kind is the destination a single extracted item was routed to.
type Kind string

const (
	KindBook       Kind = "book"        // goodreads link -> content/_data/book.json
	KindMovie      Kind = "movie"       // imdb link -> content/_data/movie.json
	KindRecipeLink Kind = "recipe-link" // recipe URL -> content/blog/recipes.md
	KindRecipeText Kind = "recipe-text" // pasted recipe body -> draft, needs manual formatting
	KindFeed       Kind = "feed"        // everything else -> draft feed block
	KindText       Kind = "text"        // prose with no URL -> draft text section
	KindSkipped    Kind = "skipped"     // filtered host -> draft, kept out of the feed
	KindRating     Kind = "rating"      // "Title 4/5" -> draft, formatted for movies.md
)

// Item is one extracted unit of material after classification and resolution.
type Item struct {
	Kind  Kind
	URL   string
	Title string
	Body  string // for text/recipe-text items
	Note  Note

	// Resolved reports that URL was derived from a different source (e.g. a
	// Wikipedia page) rather than supplied directly. Resolved items are always
	// staged for review instead of written straight into the curated lists,
	// because a search-based match can pick the wrong entry.
	Resolved bool
	From     string // the original URL a resolved item came from
	Note2    string // why it needs review

	Rating int // 1-5, for KindRating items
}

// Entry is one record in book.json / movie.json.
type Entry struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
