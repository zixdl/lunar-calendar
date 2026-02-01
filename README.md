# Lunar Calendar CLI

A command-line tool for converting Gregorian (solar) dates to Lunar calendar dates and displaying lunar month calendars. Built with Go following clean architecture principles.

## Features

- **Date Conversion**: Convert any Gregorian date to its corresponding Lunar date
- **Current Date**: Get today's Lunar date with no arguments
- **Calendar Display**: View full monthly calendars in a format similar to the Linux `cal` command
- **Date Highlighting**: Current date is highlighted in calendar view (reverse video)
- **Accurate Calculations**: Uses astronomical algorithms for precise lunar date calculations
- **Time Zone Support**: Configured for UTC+7 (Southeast Asia/Vietnam)

## Installation

### Prerequisites

- Go 1.16 or higher

### Install from Source

```bash
# Clone the repository
git clone <repository-url>
cd lunar-calendar

# Build and install
go install ./cmd/lunar
```

Or build locally:

```bash
go build -o lunar ./cmd/lunar
```

## Usage

### Convert Gregorian Date to Lunar Date

Convert a specific date:

```bash
lunar -d "2026-02-01"
```

Output:

```text
2025-12-14
```

Convert today's date (no arguments):

```bash
lunar
```

Output:

```text
2025-12-14
```

### Display Lunar Calendar

Show calendar for a specific lunar month:

```bash
lunar -cal "12"
```

Output:

```text
   December 2025
Su Mo Tu We Th Fr Sa
    1  2  3  4  5  6
 7  8  9 10 11 12 13
14 15 16 17 18 19 20   # Today's date (14) is highlighted
21 22 23 24 25 26 27
28 29
```

More examples:

```bash
# Show lunar January calendar
lunar -cal "1"

# Show lunar February calendar
lunar -cal "2"
```

Output:

```text
   January 2025
Su Mo Tu We Th Fr Sa
          1  2  3  4
 5  6  7  8  9 10 11
12 13 14 15 16 17 18
19 20 21 22 23 24 25
26 27 28 29 30
```

## Command-Line Options

| Flag   | Description                                          | Example             | Default       |
|--------|------------------------------------------------------|---------------------|---------------|
| `-d`   | Specify a Gregorian date to convert (YYYY-MM-DD)     | `-d "2026-02-01"`   | Today's date  |
| `-cal` | Display calendar for specified lunar month (1-12)    | `-cal "12"`         | N/A           |

## Building and Testing

### Build

```bash
# Build the executable
go build -o lunar ./cmd/lunar

# Build and install to $GOPATH/bin
go install ./cmd/lunar
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Limitations

- Configured for UTC+7 timezone
- Calendar search algorithm optimized for recent dates (1900-2100)
- Does not display leap month indicators

## Acknowledgments

- Lunar calendar algorithms based on astronomical calculations
- Inspired by the Unix/Linux `cal` command
- Vietnamese lunar calendar system
