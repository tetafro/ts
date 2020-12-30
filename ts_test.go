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
		ok   bool
	}{
		{
			name: "int timestamp",
			in:   "1562913751",
			out:  time.Date(2019, 7, 12, 6, 42, 31, 0, time.UTC),
			ok:   true,
		},
		{
			name: "int timestamp with milliseconds",
			in:   "1562913751555",
			out:  time.Date(2019, 7, 12, 6, 42, 31, 0, time.UTC),
			ok:   true,
		},
		{
			name: "float timestamp",
			in:   "1562913751.0",
			out:  time.Date(2019, 7, 12, 6, 42, 31, 0, time.UTC),
			ok:   true,
		},
		{
			name: "float timestamp with milliseconds",
			in:   "1562913751555.0",
			out:  time.Date(2019, 7, 12, 6, 42, 31, 0, time.UTC),
			ok:   true,
		},
		{
			name: "date",
			in:   "2019-01-01",
			out:  time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
			ok:   true,
		},
		{
			name: "plus hours",
			in:   "+1h",
			out:  time.Date(2019, 10, 10, 1, 0, 0, 0, time.UTC),
			ok:   true,
		},
		{
			name: "minus hours",
			in:   "-1h",
			out:  time.Date(2019, 10, 9, 23, 0, 0, 0, time.UTC),
			ok:   true,
		},
		{
			name: "plus days",
			in:   "+1d",
			out:  time.Date(2019, 10, 11, 0, 0, 0, 0, time.UTC),
			ok:   true,
		},
		{
			name: "minus days",
			in:   "-1d",
			out:  time.Date(2019, 10, 9, 0, 0, 0, 0, time.UTC),
			ok:   true,
		},
		{
			name: "date, time",
			in:   "2019/05/23 12:54",
			out:  time.Date(2019, 5, 23, 12, 54, 0, 0, time.UTC),
			ok:   true,
		},
		{
			name: "RFC 3339",
			in:   "2019-07-15T15:00:10+00:00",
			out:  time.Date(2019, 7, 15, 15, 0, 10, 0, time.UTC),
			ok:   true,
		},
		{
			name: "RFC 3339",
			in:   "2019-07-15T15:00:10+03:00",
			out:  time.Date(2019, 7, 15, 12, 0, 10, 0, time.UTC),
			ok:   true,
		},
		{
			name: "invalid",
			in:   "hello",
			out:  time.Time{},
			ok:   false,
		},
	}

	for _, tt := range testCases {
		now := time.Date(2019, 10, 10, 0, 0, 0, 0, time.UTC)
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			out, ok := parse(tt.in, now)
			if tt.ok != ok {
				t.Fatalf("Wrong result\n  expected: %v\n       got: %v", tt.ok, ok)
			}
			if tt.out != out {
				t.Fatalf("Wrong time\n  expected: %s\n       got: %s", tt.out, out)
			}
		})
	}
}
