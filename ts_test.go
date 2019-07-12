package main

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  time.Time
	}{
		{
			name: "timestamp",
			in:   "1562913751",
			out:  time.Date(2019, 07, 12, 6, 42, 31, 0, time.UTC),
		},
		{
			name: "date",
			in:   "2019-01-01",
			out:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "date, time",
			in:   "2019/05/23 12:54",
			out:  time.Date(2019, 5, 23, 12, 54, 0, 0, time.UTC),
		},
		{
			name: "RFC 3339",
			in:   "2019-07-15T15:00:10+00:00",
			out:  time.Date(2019, 7, 15, 15, 0, 10, 0, time.UTC),
		},
		{
			name: "RFC 3339",
			in:   "2019-07-15T15:00:10+03:00",
			out:  time.Date(2019, 7, 15, 12, 0, 10, 0, time.UTC),
		},
		{
			name: "invalid",
			in:   "hello",
			out:  time.Time{},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			out := parse(tt.in)
			if tt.out != out {
				t.Fatalf("Wrong time\n  expected: %s\n       got: %s", tt.out, out)
			}
		})
	}
}
