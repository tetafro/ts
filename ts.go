package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	t := time.Now()
	if len(args) > 0 {
		in := strings.Join(args, " ")

		var ok bool
		t, ok = parse(in, t)
		if !ok {
			fmt.Println("Failed to parse input")
			os.Exit(1)
		}
	}

	show(t.Local())
}

func parse(in string, now time.Time) (time.Time, bool) {
	d, ok := parseDiff(in)
	if ok {
		return now.Add(d), true
	}

	t, ok := parseTime(in)
	if ok {
		return t, true
	}
	return time.Time{}, false
}

func parseDiff(in string) (time.Duration, bool) {
	var mult time.Duration
	if strings.HasPrefix(in, "+") {
		in = strings.TrimLeft(in, "+")
		mult = 1
	} else if strings.HasPrefix(in, "-") {
		mult = -1
	} else {
		return 0, false
	}

	in = strings.TrimLeft(in, "-")

	var days bool
	if d, err := time.ParseDuration(in); err == nil {
		return mult * d, true
	}
	if strings.HasSuffix(in, "d") {
		in = strings.TrimRight(in, "d")
		days = true
	}
	if n, err := strconv.Atoi(in); err == nil {
		if days {
			return mult * time.Duration(n) * 24 * time.Hour, true
		}
		return mult * time.Duration(n) * time.Second, true
	}
	if n, err := strconv.ParseFloat(in, 64); err == nil {
		if days {
			return mult * time.Duration(n) * 24 * time.Hour, true
		}
		return mult * time.Duration(n) * time.Second, true
	}
	return 0, false
}

func parseTime(in string) (time.Time, bool) {
	// Timestamp in seconds
	if n, err := strconv.Atoi(in); err == nil {
		if n > 10000000000 { // milliseconds
			return time.Unix(int64(n/1000), 0).UTC(), true
		}
		return time.Unix(int64(n), 0).UTC(), true
	}
	if n, err := strconv.ParseFloat(in, 64); err == nil {
		if n > 10000000000 { // milliseconds
			return time.Unix(int64(n/1000), 0).UTC(), true
		}
		return time.Unix(int64(n), 0).UTC(), true
	}

	formats := []string{
		// Reverse date
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
		// Straightforward date
		"02-01-2006",
		"02/01/2006",
		"02.01.2006",
		// Reverse datetime
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006.01.02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04",
		"2006.01.02 15:04",
		// Reverse datetime with timezone name
		"2006-01-02 15:04:05 MST",
		"2006/01/02 15:04:05 MST",
		"2006.01.02 15:04:05 MST",
		"2006-01-02 15:04 MST",
		"2006/01/02 15:04 MST",
		"2006.01.02 15:04 MST",
		// Reverse datetime with timezone shift
		"2006-01-02 15:04:05 -07:00",
		"2006/01/02 15:04:05 -07:00",
		"2006.01.02 15:04:05 -07:00",
		"2006-01-02 15:04 -07:00",
		"2006/01/02 15:04 -07:00",
		"2006.01.02 15:04 -07:00",
		// RFC
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
	}
	for _, format := range formats {
		t, err := time.Parse(format, in)
		if err == nil {
			return t.UTC(), true
		}
	}
	return time.Time{}, false
}

func show(t time.Time) {
	fmt.Printf("Timestamp: %d\n", t.Unix())
	fmt.Printf("RFC 3339:  %s\n", t.Format(time.RFC3339))
	fmt.Printf("RFC 1123:  %s\n", t.Format(time.RFC1123))
}
