package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// A Signal chat export is line based: a date header, then each message followed
// by its own HH:MM line, with blank lines between messages.
//
//	13 Mar 2026
//
//	https://sharif.io/looking-stupid
//	3:30
//
//	Fri 10 Apr
//
//	http://beigeboxoss.com/
//	14:44
var (
	// "Fri 10 Apr", "Fri, 10 Apr 2026", "10 Apr 2026", "10 April"
	dayFirstRe = regexp.MustCompile(`(?i)^(?:(?:mon|tue|tues|wed|thu|thur|thurs|fri|sat|sun)[a-z]*\.?,?\s+)?(\d{1,2})\.?\s+([a-z]{3,9})\.?(?:,?\s+(\d{4}))?$`)
	// "Mar 13, 2026", "March 13"
	monthFirstRe = regexp.MustCompile(`(?i)^(?:(?:mon|tue|tues|wed|thu|thur|thurs|fri|sat|sun)[a-z]*\.?,?\s+)?([a-z]{3,9})\.?\s+(\d{1,2})(?:,?\s+(\d{4}))?$`)
	// "13/03/2026" or "13.03.2026"
	numericDateRe = regexp.MustCompile(`^(\d{1,2})[./-](\d{1,2})[./-](\d{2,4})$`)
	relativeRe    = regexp.MustCompile(`(?i)^(today|yesterday)$`)
	timeRe        = regexp.MustCompile(`(?i)^(\d{1,2}):(\d{2})(?::\d{2})?\s*(am|pm)?$`)
)

var months = map[string]time.Month{
	"jan": time.January, "feb": time.February, "mar": time.March,
	"apr": time.April, "may": time.May, "jun": time.June,
	"jul": time.July, "aug": time.August, "sep": time.September,
	"oct": time.October, "nov": time.November, "dec": time.December,
}

func monthFromName(s string) (time.Month, bool) {
	s = strings.ToLower(strings.TrimSpace(s))
	if len(s) < 3 {
		return 0, false
	}
	m, ok := months[s[:3]]
	return m, ok
}

// parseDateHeader recognises a Signal day separator. ref is "now", used to infer
// the year when the header omits it.
func parseDateHeader(line string, ref time.Time) (time.Time, bool) {
	line = strings.TrimSpace(line)
	if line == "" {
		return time.Time{}, false
	}

	if m := relativeRe.FindStringSubmatch(line); m != nil {
		d := startOfDay(ref)
		if strings.EqualFold(m[1], "yesterday") {
			d = d.AddDate(0, 0, -1)
		}
		return d, true
	}

	if m := numericDateRe.FindStringSubmatch(line); m != nil {
		day, _ := strconv.Atoi(m[1])
		mon, _ := strconv.Atoi(m[2])
		year, _ := strconv.Atoi(m[3])
		if year < 100 {
			year += 2000
		}
		if mon < 1 || mon > 12 {
			return time.Time{}, false
		}
		return buildDate(year, time.Month(mon), day, ref, true)
	}

	if m := dayFirstRe.FindStringSubmatch(line); m != nil {
		if mon, ok := monthFromName(m[2]); ok {
			day, _ := strconv.Atoi(m[1])
			return fromParts(day, mon, m[3], ref)
		}
	}

	if m := monthFirstRe.FindStringSubmatch(line); m != nil {
		if mon, ok := monthFromName(m[1]); ok {
			day, _ := strconv.Atoi(m[2])
			return fromParts(day, mon, m[3], ref)
		}
	}

	return time.Time{}, false
}

func fromParts(day int, mon time.Month, yearStr string, ref time.Time) (time.Time, bool) {
	if yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			return time.Time{}, false
		}
		return buildDate(year, mon, day, ref, true)
	}
	return buildDate(ref.Year(), mon, day, ref, false)
}

// buildDate validates the date. When the year was not stated and the result
// would fall in the future, it belongs to the previous year instead -- notes are
// always about things already posted.
func buildDate(year int, mon time.Month, day int, ref time.Time, yearKnown bool) (time.Time, bool) {
	if day < 1 || day > 31 {
		return time.Time{}, false
	}
	d := time.Date(year, mon, day, 0, 0, 0, 0, time.Local)
	if d.Day() != day || d.Month() != mon {
		return time.Time{}, false // e.g. 31 Feb
	}
	if !yearKnown && d.After(startOfDay(ref)) {
		d = d.AddDate(-1, 0, 0)
	}
	return d, true
}

func startOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// parseTimeOfDay recognises the HH:MM line Signal prints under each message.
func parseTimeOfDay(line string) (time.Duration, bool) {
	m := timeRe.FindStringSubmatch(strings.TrimSpace(line))
	if m == nil {
		return 0, false
	}
	h, err := strconv.Atoi(m[1])
	if err != nil {
		return 0, false
	}
	min, err := strconv.Atoi(m[2])
	if err != nil || min > 59 {
		return 0, false
	}
	switch strings.ToLower(m[3]) {
	case "pm":
		if h < 12 {
			h += 12
		}
	case "am":
		if h == 12 {
			h = 0
		}
	}
	if h > 23 {
		return 0, false
	}
	return time.Duration(h)*time.Hour + time.Duration(min)*time.Minute, true
}
