// Copyright (c) 2015-2024 libreFS, Inc.
//
// This file is part of libreFS Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"testing"
	"time"
)

// ── humanizedDuration.String ──────────────────────────────────────────────────

func TestHumanizedDurationString(t *testing.T) {
	cases := []struct {
		d    humanizedDuration
		want string
	}{
		{humanizedDuration{MilliSeconds: 500}, "500 milliseconds"},
		{humanizedDuration{Seconds: 45}, "45 seconds"},
		{humanizedDuration{Minutes: 3, Seconds: 20}, "3 minutes 20 seconds"},
		{humanizedDuration{Hours: 2, Minutes: 15, Seconds: 5}, "2 hours 15 minutes 5 seconds"},
		{humanizedDuration{Days: 1, Hours: 3, Minutes: 10, Seconds: 0}, "1 days 3 hours 10 minutes 0 seconds"},
	}
	for _, c := range cases {
		got := c.d.String()
		if got != c.want {
			t.Errorf("String(%+v) = %q, want %q", c.d, got, c.want)
		}
	}
}

// ── humanizedDuration.StringShort ────────────────────────────────────────────

func TestHumanizedDurationStringShort(t *testing.T) {
	cases := []struct {
		d    humanizedDuration
		want string
	}{
		// only milliseconds
		{humanizedDuration{MilliSeconds: 250}, "250 milliseconds"},
		// only seconds (no minutes/hours/days)
		{humanizedDuration{Seconds: 30}, "30 seconds"},
		// minutes only (no hours/days)
		{humanizedDuration{Minutes: 5, Seconds: 10}, "5 minutes"},
		// hours + minutes
		{humanizedDuration{Hours: 2, Minutes: 30}, "2 hours 30 minutes"},
		// 1 day (≤2) → days + hours
		{humanizedDuration{Days: 1, Hours: 6}, "1 days, 6 hours"},
		// 2 days (≤2) → days + hours
		{humanizedDuration{Days: 2, Hours: 0}, "2 days, 0 hours"},
		// >2 days → just days
		{humanizedDuration{Days: 10, Hours: 3}, "10 days"},
	}
	for _, c := range cases {
		got := c.d.StringShort()
		if got != c.want {
			t.Errorf("StringShort(%+v) = %q, want %q", c.d, got, c.want)
		}
	}
}

// ── timeDurationToHumanizedDuration ──────────────────────────────────────────

func TestTimeDurationToHumanizedDuration(t *testing.T) {
	cases := []struct {
		in   time.Duration
		want humanizedDuration
	}{
		// < 1 second → milliseconds
		{500 * time.Millisecond, humanizedDuration{MilliSeconds: 500}},
		// < 1 minute → seconds only
		{45 * time.Second, humanizedDuration{Seconds: 45}},
		// < 1 hour → minutes + seconds
		{3*time.Minute + 20*time.Second, humanizedDuration{Minutes: 3, Seconds: 20}},
		// < 1 day → hours + minutes + seconds
		{2*time.Hour + 15*time.Minute + 5*time.Second, humanizedDuration{Hours: 2, Minutes: 15, Seconds: 5}},
		// ≥ 1 day → days + hours + minutes + seconds
		{25*time.Hour + 10*time.Minute + 30*time.Second, humanizedDuration{Days: 1, Hours: 1, Minutes: 10, Seconds: 30}},
	}
	for _, c := range cases {
		got := timeDurationToHumanizedDuration(c.in)
		if got != c.want {
			t.Errorf("timeDurationToHumanizedDuration(%v) = %+v, want %+v", c.in, got, c.want)
		}
	}
}
