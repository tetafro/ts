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
		t = parse(in)
	}

	if t.IsZero() {
		fmt.Println("Failed to parse input")
		os.Exit(1)
	}
	show(t)
}

func parse(in string) time.Time {
	// Timestamp in seconds
	if n, err := strconv.Atoi(in); err == nil {
		return time.Unix(int64(n), 0).UTC()
	}
	if n, err := strconv.ParseFloat(in, 64); err == nil {
		return time.Unix(int64(n), 0).UTC()
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
			return t.UTC()
		}
	}
	return time.Time{}
}

func show(t time.Time) {
	fmt.Printf("Timestamp: %d\n", t.Unix())
	fmt.Printf("RFC 3339:  %s\n", t.Format(time.RFC3339))
	fmt.Printf("RFC 1123:  %s\n", t.Format(time.RFC1123))
}
