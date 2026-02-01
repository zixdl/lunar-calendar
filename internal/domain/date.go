package domain

import "time"

// LunarDate represents a date in the lunar calendar
type LunarDate struct {
	Year  int
	Month int
	Day   int
}

// SolarDate represents a date in the Gregorian (solar) calendar
type SolarDate struct {
	Year  int
	Month int
	Day   int
}

// FromTime creates a SolarDate from a time.Time
func FromTime(t time.Time) SolarDate {
	return SolarDate{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}

// ToTime converts a SolarDate to time.Time
func (s SolarDate) ToTime() time.Time {
	return time.Date(s.Year, time.Month(s.Month), s.Day, 0, 0, 0, 0, time.UTC)
}
