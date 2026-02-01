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

## Project Structure

```text
lunar-calendar/
├── cmd/
│   └── lunar/
│       └── main.go              # Application entry point
├── internal/
│   ├── domain/
│   │   ├── date.go             # Domain models (LunarDate, SolarDate)
│   │   └── date_test.go        # Domain tests
│   ├── converter/
│   │   ├── converter.go        # Lunar calendar conversion algorithms
│   │   └── converter_test.go   # Converter tests
│   ├── calendar/
│   │   └── calendar.go         # Calendar display logic
│   └── cli/
│       └── cli.go              # CLI interface and flag parsing
├── go.mod
└── README.md
```

## Architecture

This project follows **Clean Architecture** principles:

### Layers

1. **Domain Layer** (`internal/domain`)
   - Contains core business entities (LunarDate, SolarDate)
   - No external dependencies
   - Pure data structures

2. **Use Case Layer** (`internal/converter`, `internal/calendar`)
   - `converter`: Implements lunar calendar conversion algorithms
   - `calendar`: Handles calendar display logic
   - Business logic and calculations

3. **Interface Layer** (`internal/cli`)
   - CLI interface implementation
   - Flag parsing and user interaction
   - Coordinates between use cases

4. **Main** (`cmd/lunar`)
   - Application entry point
   - Dependency injection

### Key Design Principles

- **Separation of Concerns**: Each package has a single, well-defined responsibility
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Testability**: Core logic is easily testable without external dependencies
- **Encapsulation**: Internal implementation details are hidden

## Algorithm

The lunar calendar conversion uses astronomical calculations based on:

1. **Julian Day Number**: Converting Gregorian dates to Julian day numbers
2. **New Moon Calculation**: Determining lunar month starts using astronomical formulas
3. **Sun Longitude**: Calculating solar terms for accurate month determination
4. **Leap Month Detection**: Identifying intercalary months in the lunar calendar

The implementation is based on the Vietnamese lunar calendar system, which follows the Chinese lunar calendar with timezone adjustments for UTC+7.

### Key Functions

- `jdFromDate()`: Converts Gregorian date to Julian day number
- `getNewMoonDay()`: Calculates the date of a new moon
- `getSunLongitude()`: Determines the sun's position
- `getLunarMonth11()`: Finds the 11th lunar month of a year
- `getLeapMonthOffset()`: Identifies leap month positions

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

### Test Coverage

The project includes unit tests for:

- Domain models (date conversions)
- Converter algorithms (Julian day calculations, lunar conversions)
- Edge cases and boundary conditions

Example test output:

```text
ok      lunar-calendar/internal/converter    0.446s
ok      lunar-calendar/internal/domain       0.269s
```

## Technical Details

### Date Format

- **Input**: `YYYY-MM-DD` (e.g., `2026-02-01`)
- **Output**: `YYYY-MM-DD` (e.g., `2025-12-14`)

### Calendar Display Features

- **Week Start**: Sunday
- **Date Highlighting**: Current date shown in reverse video (ANSI escape codes)
- **Alignment**: Days are right-aligned in 3-character columns
- **Format**: Similar to Unix/Linux `cal` command

### Time Zone

The converter is configured for **UTC+7** (Southeast Asia), suitable for:

- Vietnam
- Thailand
- Cambodia
- Laos
- Western Indonesia

## Development

### Code Style

- Follows Go best practices and idioms
- Uses `gofmt` for code formatting
- Adheres to clean code principles
- Comprehensive comments for complex algorithms

### Requirements

- Go 1.16+
- No external dependencies (uses only Go standard library)

## Examples

### Example 1: Check Today's Lunar Date

```bash
$ lunar
2025-12-14
```

### Example 2: Historical Date Conversion

```bash
$ lunar -d "2000-01-01"
1999-11-25
```

### Example 3: View Current Month Calendar

```bash
$ lunar -cal "12"
   December 2025
Su Mo Tu We Th Fr Sa
    1  2  3  4  5  6
 7  8  9 10 11 12 13
14 15 16 17 18 19 20
21 22 23 24 25 26 27
28 29
```

## Limitations

- Configured for UTC+7 timezone
- Calendar search algorithm optimized for recent dates (1900-2100)
- Does not display leap month indicators

## Contributing

Contributions are welcome! Please ensure:

1. Code follows Go best practices
2. All tests pass (`go test ./...`)
3. New features include tests
4. Documentation is updated

## License

[Specify your license here]

## Acknowledgments

- Lunar calendar algorithms based on astronomical calculations
- Inspired by the Unix/Linux `cal` command
- Vietnamese lunar calendar system
