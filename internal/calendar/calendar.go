package calendar

import (
	"fmt"
	"lunar-calendar/internal/converter"
	"lunar-calendar/internal/domain"
	"strings"
	"time"
)

// Calendar handles calendar display operations
type Calendar struct {
	converter *converter.Converter
}

// NewCalendar creates a new calendar instance
func NewCalendar(conv *converter.Converter) *Calendar {
	return &Calendar{
		converter: conv,
	}
}

// GetMonthCalendar returns a formatted calendar for the specified lunar month
func (c *Calendar) GetMonthCalendar(lunarMonth, lunarYear int) string {
	var result strings.Builder

	// Find the solar dates corresponding to this lunar month
	solarDates := c.findSolarDatesForLunarMonth(lunarMonth, lunarYear)

	if len(solarDates) == 0 {
		return fmt.Sprintf("No dates found for lunar month %d, year %d\n", lunarMonth, lunarYear)
	}

	// Build a map of lunar day -> solar date
	lunarDayMap := make(map[int]domain.SolarDate)
	for _, solar := range solarDates {
		lunar := c.converter.ToLunar(solar)
		if lunar.Month == lunarMonth && lunar.Year == lunarYear {
			lunarDayMap[lunar.Day] = solar
		}
	}

	// Get today's lunar date for highlighting
	today := time.Now()
	todayLunar := c.converter.ToLunar(domain.FromTime(today))
	todayDay := -1
	if todayLunar.Month == lunarMonth && todayLunar.Year == lunarYear {
		todayDay = todayLunar.Day
	}

	// Get the month name
	monthName := c.getLunarMonthName(lunarMonth)
	header := fmt.Sprintf("   %s %d", monthName, lunarYear)

	// Center the header
	result.WriteString(header + "\n")

	// Day of week headers
	result.WriteString("Su Mo Tu We Th Fr Sa\n")

	// Find the first day and its weekday
	if len(lunarDayMap) == 0 {
		return result.String()
	}

	firstSolar := lunarDayMap[1]
	firstWeekday := int(firstSolar.ToTime().Weekday())

	// Add leading spaces for days before the first day
	for i := 0; i < firstWeekday; i++ {
		result.WriteString("   ")
	}

	// Print the days
	currentWeekday := firstWeekday
	maxDay := 0
	for day := range lunarDayMap {
		if day > maxDay {
			maxDay = day
		}
	}

	for day := 1; day <= maxDay; day++ {
		if _, exists := lunarDayMap[day]; exists {
			if day == todayDay {
				// Highlight today's date with reverse video
				result.WriteString(fmt.Sprintf("\033[7m%2d\033[0m ", day))
			} else {
				result.WriteString(fmt.Sprintf("%2d ", day))
			}
		} else {
			result.WriteString("   ")
		}

		currentWeekday++
		if currentWeekday > 6 {
			result.WriteString("\n")
			currentWeekday = 0
		}
	}

	// Add final newline if needed
	if currentWeekday != 0 {
		result.WriteString("\n")
	}

	return result.String()
}

// findSolarDatesForLunarMonth finds solar dates that correspond to a lunar month
func (c *Calendar) findSolarDatesForLunarMonth(lunarMonth, lunarYear int) []domain.SolarDate {
	var dates []domain.SolarDate

	// Start from an appropriate solar year and month
	// Lunar month 1 typically starts in Jan-Feb of the solar year
	// Lunar month 12 typically starts in Dec-Jan of the next solar year
	startSolarYear := lunarYear
	startSolarMonth := 1

	// For later lunar months, start search from later in the year
	if lunarMonth >= 11 {
		startSolarMonth = 11
	} else if lunarMonth >= 7 {
		startSolarMonth = 7
	} else if lunarMonth >= 4 {
		startSolarMonth = 4
	}

	startDate := time.Date(startSolarYear, time.Month(startSolarMonth), 1, 0, 0, 0, 0, time.UTC)

	// Search through up to 120 days to find all days of the lunar month
	consecutiveNonMatches := 0
	for i := 0; i < 120; i++ {
		currentDate := startDate.AddDate(0, 0, i)
		solar := domain.FromTime(currentDate)
		lunar := c.converter.ToLunar(solar)

		if lunar.Month == lunarMonth && lunar.Year == lunarYear {
			dates = append(dates, solar)
			consecutiveNonMatches = 0
		} else if len(dates) > 0 {
			consecutiveNonMatches++
			// Stop if we've found dates and then had 5 consecutive non-matches
			if consecutiveNonMatches >= 5 {
				break
			}
		}
	}

	return dates
}

// getLunarMonthName returns the name of the lunar month
func (c *Calendar) getLunarMonthName(month int) string {
	monthNames := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	if name, ok := monthNames[month]; ok {
		return name
	}
	return fmt.Sprintf("Month %d", month)
}

// getDayOfWeek returns the abbreviated day name for a solar date
func (c *Calendar) getDayOfWeek(solar domain.SolarDate) string {
	t := solar.ToTime()
	return t.Weekday().String()[:3]
}
