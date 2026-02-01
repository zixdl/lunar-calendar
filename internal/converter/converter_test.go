package converter

import (
	"lunar-calendar/internal/domain"
	"testing"
)

func TestToLunar(t *testing.T) {
	conv := NewConverter()

	tests := []struct {
		name     string
		solar    domain.SolarDate
		expected domain.LunarDate
	}{
		{
			name: "2026-02-01 should convert to 2025-12-14",
			solar: domain.SolarDate{
				Year:  2026,
				Month: 2,
				Day:   1,
			},
			expected: domain.LunarDate{
				Year:  2025,
				Month: 12,
				Day:   14,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := conv.ToLunar(tt.solar)

			if result.Year != tt.expected.Year {
				t.Errorf("Year: got %d, want %d", result.Year, tt.expected.Year)
			}
			if result.Month != tt.expected.Month {
				t.Errorf("Month: got %d, want %d", result.Month, tt.expected.Month)
			}
			if result.Day != tt.expected.Day {
				t.Errorf("Day: got %d, want %d", result.Day, tt.expected.Day)
			}
		})
	}
}

func TestJdFromDate(t *testing.T) {
	conv := NewConverter()

	tests := []struct {
		name     string
		day      int
		month    int
		year     int
		expected int
	}{
		{
			name:     "2000-01-01",
			day:      1,
			month:    1,
			year:     2000,
			expected: 2451545,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := conv.jdFromDate(tt.day, tt.month, tt.year)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
