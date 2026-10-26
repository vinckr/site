package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

// signalEnvelope is the subset of `signal-cli -o json receive` output we care about.
type signalEnvelope struct {
	Envelope struct {
		Timestamp    int64  `json:"timestamp"`
		SourceNumber string `json:"sourceNumber"`
		DataMessage  *struct {
			Message string `json:"message"`
		} `json:"dataMessage"`
		SyncMessage *struct {
			SentMessage *struct {
				Timestamp         int64  `json:"timestamp"`
				DestinationNumber string `json:"destinationNumber"`
				Message           string `json:"message"`
			} `json:"sentMessage"`
		} `json:"syncMessage"`
	} `json:"envelope"`
}

// drainSignal runs signal-cli once and appends any new Note-to-Self messages to
// the local store. signal-cli delivers each message only once, so the store is
// the durable record -- draft processing always reads from the store, never
// directly from signal-cli.
func drainSignal(account string) ([]Note, error) {
	if account == "" {
		return nil, fmt.Errorf("no Signal account set (SIGNAL_ACCOUNT in .env or -account)")
	}
	if _, err := exec.LookPath("signal-cli"); err != nil {
		return nil, fmt.Errorf("signal-cli not installed: %w", err)
	}

	cmd := exec.Command("signal-cli", "-a", account, "-o", "json", "receive")
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("signal-cli receive: %v: %s", err, strings.TrimSpace(string(ee.Stderr)))
		}
		return nil, fmt.Errorf("signal-cli receive: %w", err)
	}

	var notes []Note
	sc := bufio.NewScanner(strings.NewReader(string(out)))
	sc.Buffer(make([]byte, 0, 1024*1024), 8*1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		var env signalEnvelope
		if err := json.Unmarshal([]byte(line), &env); err != nil {
			continue // ignore non-JSON diagnostics on stdout
		}
		e := env.Envelope

		// Note to Self arrives as a sync message addressed to your own number.
		if sm := e.SyncMessage; sm != nil && sm.SentMessage != nil {
			s := sm.SentMessage
			if s.DestinationNumber == account && strings.TrimSpace(s.Message) != "" {
				ts := s.Timestamp
				if ts == 0 {
					ts = e.Timestamp
				}
				notes = append(notes, Note{Timestamp: ts, Text: s.Message, Source: "signal"})
			}
			continue
		}
		// A message sent from this linked device to itself.
		if dm := e.DataMessage; dm != nil && e.SourceNumber == account && strings.TrimSpace(dm.Message) != "" {
			notes = append(notes, Note{Timestamp: e.Timestamp, Text: dm.Message, Source: "signal"})
		}
	}
	return notes, sc.Err()
}

// readFileNotes parses a Signal chat export (or any plain list of links). Date
// headers set the day for everything that follows, and the HH:MM line Signal
// prints under each message both closes that message and orders it within the
// day. Lines that are neither become message content.
func readFileNotes(path string) ([]Note, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ref := time.Now()
	day := startOfDay(ref)
	var (
		notes []Note
		buf   []string
		seq   int
	)

	// flush turns the buffered lines into one note stamped with the current day
	// and, when known, the time printed beneath the message.
	flush := func(tod time.Duration, haveTime bool) {
		text := strings.TrimSpace(strings.Join(buf, "\n"))
		buf = buf[:0]
		if text == "" {
			return
		}
		ts := day
		if haveTime {
			ts = day.Add(tod)
		}
		// A few milliseconds of sequence keeps same-minute messages ordered and
		// their store keys distinct.
		notes = append(notes, Note{
			Timestamp: ts.UnixMilli() + int64(seq),
			Text:      text,
			Source:    "file",
		})
		seq++
	}

	for _, line := range strings.Split(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\n") {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			flush(0, false)
			continue
		}
		if d, ok := parseDateHeader(trimmed, ref); ok {
			flush(0, false)
			day = d
			seq = 0
			continue
		}
		if tod, ok := parseTimeOfDay(trimmed); ok {
			flush(tod, true) // the time line closes the message above it
			continue
		}
		buf = append(buf, trimmed)
	}
	flush(0, false)

	return notes, nil
}

// loadStore reads the accumulated note store.
func loadStore(path string) ([]Note, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var notes []Note
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 1024*1024), 8*1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		var n Note
		if err := json.Unmarshal([]byte(line), &n); err != nil {
			continue
		}
		notes = append(notes, n)
	}
	return notes, sc.Err()
}

// appendStore adds notes not already present (keyed by timestamp+text) and
// returns how many were new.
func appendStore(path string, existing, incoming []Note) (int, error) {
	seen := map[string]bool{}
	for _, n := range existing {
		seen[storeKey(n)] = true
	}
	var fresh []Note
	for _, n := range incoming {
		k := storeKey(n)
		if seen[k] {
			continue
		}
		seen[k] = true
		fresh = append(fresh, n)
	}
	if len(fresh) == 0 {
		return 0, nil
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	for _, n := range fresh {
		if err := enc.Encode(n); err != nil {
			return 0, err
		}
	}
	return len(fresh), nil
}

// storeKey identifies a note for deduplication. File-sourced notes are keyed on
// their text alone, because re-supplying the same file gives every line a fresh
// timestamp and would otherwise reprocess material that was already handled.
// Signal notes keep the timestamp in the key, since posting the same link again
// on a later day is a deliberate act worth capturing twice.
func storeKey(n Note) string {
	text := strings.TrimSpace(n.Text)
	if n.Source == "file" {
		return "file\x00" + text
	}
	return fmt.Sprintf("%d\x00%s", n.Timestamp, text)
}

func sortNotes(notes []Note) {
	sort.SliceStable(notes, func(i, j int) bool { return notes[i].Timestamp < notes[j].Timestamp })
}
