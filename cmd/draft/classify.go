package main

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var urlRe = regexp.MustCompile(`https?://[^\s<>"'\x60]+`)

// extractURLs pulls URLs out of a note and returns them plus the leftover prose.
// Parentheses are kept when balanced, because Wikipedia URLs such as
// .../Stalker_(1979_film) depend on them; a trailing unbalanced ")" is treated
// as sentence punctuation instead.
func extractURLs(text string) (urls []string, rest string) {
	for _, u := range urlRe.FindAllString(text, -1) {
		urls = append(urls, trimURL(u))
	}
	rest = strings.TrimSpace(urlRe.ReplaceAllString(text, ""))
	return urls, rest
}

func trimURL(u string) string {
	for {
		trimmed := strings.TrimRight(u, ".,;:!?")
		if strings.HasSuffix(trimmed, ")") && strings.Count(trimmed, ")") > strings.Count(trimmed, "(") {
			trimmed = trimmed[:len(trimmed)-1]
		}
		if trimmed == u {
			return u
		}
		u = trimmed
	}
}

func host(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	return strings.TrimPrefix(strings.ToLower(u.Hostname()), "www.")
}

var (
	imdbTitleRe = regexp.MustCompile(`imdb\.com/title/(tt\d+)`)
	grBookRe    = regexp.MustCompile(`goodreads\.com/book/show/(\d+)`)
	grOtherRe   = regexp.MustCompile(`goodreads\.com/(list|author|series)/show/(\d+)`)
	wikipediaRe = regexp.MustCompile(`^([a-z]+)\.(m\.)?wikipedia\.org$`)
)

// normalizeIMDb rewrites any IMDb title URL to the canonical stored form.
func normalizeIMDb(raw string) (string, bool) {
	m := imdbTitleRe.FindStringSubmatch(raw)
	if m == nil {
		return "", false
	}
	return "https://www.imdb.com/title/" + m[1] + "/", true
}

// normalizeGoodreads rewrites a Goodreads URL to the canonical stored form,
// stripping the title slug and query string.
func normalizeGoodreads(raw string) (string, bool) {
	if m := grBookRe.FindStringSubmatch(raw); m != nil {
		return "https://www.goodreads.com/book/show/" + m[1], true
	}
	if m := grOtherRe.FindStringSubmatch(raw); m != nil {
		return "https://www.goodreads.com/" + m[1] + "/show/" + m[2], true
	}
	return "", false
}

// skipHosts are sites whose links are kept out of the feed block. They are not
// discarded -- they go to a separate section of the draft, because "most of the
// time" is not "always" and the occasional one is worth rescuing by hand.
var skipHosts = map[string]bool{
	"instagram.com": true,
	"instagr.am":    true,
	"ig.me":         true,
	"bandcamp.com":  true, // tracked separately
}

func isSkippedHost(raw string) bool {
	h := host(raw)
	if skipHosts[h] {
		return true
	}
	// Catch regional and subdomain variants such as de.instagram.com.
	for skip := range skipHosts {
		if strings.HasSuffix(h, "."+skip) {
			return true
		}
	}
	return false
}

// recipeHosts are sites whose links are filed under recipes rather than the feed.
var recipeHosts = map[string]bool{
	"allrecipes.com": true, "seriouseats.com": true, "bonappetit.com": true,
	"epicurious.com": true, "food52.com": true, "thekitchn.com": true,
	"cooking.nytimes.com": true, "chefkoch.de": true, "smittenkitchen.com": true,
	"budgetbytes.com": true, "justonecookbook.com": true, "kingarthurbaking.com": true,
	"seriouseats.co.uk": true, "americastestkitchen.com": true, "bbcgoodfood.com": true,
	"maangchi.com": true, "woksoflife.com": true, "seriouseat.com": true,
	"rarecooking.com": true, "supercook.com": true, "tasteatlas.com": true,
}

func isRecipeURL(raw string) bool {
	if recipeHosts[host(raw)] {
		return true
	}
	p := strings.ToLower(raw)
	return strings.Contains(p, "/recipe/") || strings.Contains(p, "/recipes/")
}

// looksLikeRecipeText reports whether a prose block is a pasted recipe. It wants
// an ingredients cue plus a method cue, which keeps ordinary notes that merely
// mention food out of the recipe bucket.
func looksLikeRecipeText(s string) bool {
	l := strings.ToLower(s)
	hasIngredients := strings.Contains(l, "ingredient") ||
		strings.Contains(l, "zutaten") ||
		strings.Contains(l, "you will need")
	hasMethod := strings.Contains(l, "step") ||
		strings.Contains(l, "instructions") ||
		strings.Contains(l, "method") ||
		strings.Contains(l, "preparation") ||
		strings.Contains(l, "zubereitung") ||
		regexp.MustCompile(`(?m)^\s*\d+[.)]\s`).MatchString(s)
	return hasIngredients && hasMethod
}

func isWikipedia(raw string) bool {
	return wikipediaRe.MatchString(host(raw))
}

// ratingRe matches a note like "Horrible Bosses 3/5" -- a title followed by a
// score out of five, which is how reviews are written in movies.md.
var ratingRe = regexp.MustCompile(`^(.*\S)\s+([1-5])\s*/\s*5\.?$`)

// parseRating splits a review note into its title and score. It requires a
// title of at least two characters so a bare "3/5" is not mistaken for one.
func parseRating(text string) (title string, score int, ok bool) {
	text = strings.TrimSpace(text)
	if strings.Contains(text, "\n") {
		return "", 0, false // a review is a single line
	}
	m := ratingRe.FindStringSubmatch(text)
	if m == nil {
		return "", 0, false
	}
	title = strings.TrimSpace(m[1])
	if len([]rune(title)) < 2 {
		return "", 0, false
	}
	score, err := strconv.Atoi(m[2])
	if err != nil {
		return "", 0, false
	}
	return title, score, true
}
