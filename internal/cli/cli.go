package cli

import (
	"flag"
	"fmt"
	"lunar-calendar/internal/calendar"
	"lunar-calendar/internal/converter"
	"lunar-calendar/internal/domain"
	"os"
	"strconv"
	"time"
)

// CLI handles command-line interface operations
type CLI struct {
	converter *converter.Converter
	calendar  *calendar.Calendar
}

// NewCLI creates a new CLI instance
func NewCLI() *CLI {
	conv := converter.NewConverter()
	cal := calendar.NewCalendar(conv)

	return &CLI{
		converter: conv,
		calendar:  cal,
	}
}

// Run executes the CLI application
func (c *CLI) Run() {
	dateFlag := flag.String("d", "", "Solar date in format YYYY-MM-DD (optional, defaults to today)")
	calFlag := flag.String("cal", "", "Lunar month number (1-12) to display calendar (optional, defaults to current lunar month)")

	flag.Parse()

	// If calendar flag is set, display the calendar
	if *calFlag != "" {
		c.handleCalendar(*calFlag)
		return
	}

	// Otherwise, display the lunar date
	c.handleDate(*dateFlag)
}

// handleDate processes the date conversion request
func (c *CLI) handleDate(dateStr string) {
	var solar domain.SolarDate

	if dateStr == "" {
		// Use today's date
		now := time.Now()
		solar = domain.FromTime(now)
	} else {
		// Parse the provided date
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid date format. Please use YYYY-MM-DD\n")
			os.Exit(1)
		}
		solar = domain.FromTime(t)
	}

	// Convert to lunar
	lunar := c.converter.ToLunar(solar)

	// Output in the format specified: YYYY-MM-DD
	fmt.Printf("%04d-%02d-%02d\n", lunar.Year, lunar.Month, lunar.Day)
}

// handleCalendar processes the calendar display request
func (c *CLI) handleCalendar(monthStr string) {
	var lunarMonth int
	var lunarYear int

	if monthStr == "" {
		// Use current lunar month
		now := time.Now()
		solar := domain.FromTime(now)
		lunar := c.converter.ToLunar(solar)
		lunarMonth = lunar.Month
		lunarYear = lunar.Year
	} else {
		// Parse the provided month
		month, err := strconv.Atoi(monthStr)
		if err != nil || month < 1 || month > 12 {
			fmt.Fprintf(os.Stderr, "Error: Invalid month number. Please provide a number between 1 and 12\n")
			os.Exit(1)
		}
		lunarMonth = month

		// Use current year's lunar year
		now := time.Now()
		solar := domain.FromTime(now)
		lunar := c.converter.ToLunar(solar)
		lunarYear = lunar.Year
	}

	// Display the calendar
	calendarOutput := c.calendar.GetMonthCalendar(lunarMonth, lunarYear)
	fmt.Print(calendarOutput)
}
