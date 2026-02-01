package domain

import (
	"testing"
	"time"
)

func TestFromTime(t *testing.T) {
	testTime := time.Date(2026, 2, 1, 10, 30, 0, 0, time.UTC)
	solar := FromTime(testTime)

	if solar.Year != 2026 {
		t.Errorf("Year: got %d, want 2026", solar.Year)
	}
	if solar.Month != 2 {
		t.Errorf("Month: got %d, want 2", solar.Month)
	}
	if solar.Day != 1 {
		t.Errorf("Day: got %d, want 1", solar.Day)
	}
}

func TestToTime(t *testing.T) {
	solar := SolarDate{
		Year:  2026,
		Month: 2,
		Day:   1,
	}

	result := solar.ToTime()
	expected := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)

	if !result.Equal(expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}
