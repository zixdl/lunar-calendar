package converter

import (
	"lunar-calendar/internal/domain"
	"math"
)

const timeZone = 7.0 // UTC+7 for Vietnam/Southeast Asia

// Converter handles conversion between solar and lunar calendars
type Converter struct{}

// NewConverter creates a new converter instance
func NewConverter() *Converter {
	return &Converter{}
}

// ToLunar converts a solar date to lunar date
func (c *Converter) ToLunar(solar domain.SolarDate) domain.LunarDate {
	dayNumber := c.jdFromDate(solar.Day, solar.Month, solar.Year)
	k := int((float64(dayNumber) - 2415021.076998695) / 29.530588853)
	monthStart := c.getNewMoonDay(k+1)

	if monthStart > dayNumber {
		monthStart = c.getNewMoonDay(k)
	}

	a11 := c.getLunarMonth11(solar.Year)
	b11 := a11
	lunarYear := solar.Year

	if a11 >= monthStart {
		lunarYear = solar.Year
		a11 = c.getLunarMonth11(solar.Year - 1)
	} else {
		lunarYear = solar.Year + 1
		b11 = c.getLunarMonth11(solar.Year + 1)
	}

	lunarDay := dayNumber - monthStart + 1
	diff := int((float64(monthStart) - float64(a11)) / 29)
	lunarMonth := diff + 11

	if b11-a11 > 365 {
		leapMonthDiff := c.getLeapMonthOffset(a11)
		if diff >= leapMonthDiff {
			lunarMonth = diff + 10
		}
	}

	if lunarMonth > 12 {
		lunarMonth = lunarMonth - 12
	}
	if lunarMonth >= 11 && diff < 4 {
		lunarYear -= 1
	}

	return domain.LunarDate{
		Year:  lunarYear,
		Month: lunarMonth,
		Day:   int(lunarDay),
	}
}

// jdFromDate calculates Julian day number from day, month, year
func (c *Converter) jdFromDate(dd, mm, yy int) int {
	a := (14 - mm) / 12
	y := yy + 4800 - a
	m := mm + 12*a - 3
	jd := dd + (153*m+2)/5 + 365*y + y/4 - y/100 + y/400 - 32045
	return jd
}

// getNewMoonDay calculates the Julian day number of the k-th new moon
func (c *Converter) getNewMoonDay(k int) int {
	T := float64(k) / 1236.85
	T2 := T * T
	T3 := T2 * T
	dr := math.Pi / 180.0
	Jd1 := 2415020.75933 + 29.53058868*float64(k) + 0.0001178*T2 - 0.000000155*T3
	Jd1 = Jd1 + 0.00033*math.Sin((166.56+132.87*T-0.009173*T2)*dr)
	M := 359.2242 + 29.10535608*float64(k) - 0.0000333*T2 - 0.00000347*T3
	Mpr := 306.0253 + 385.81691806*float64(k) + 0.0107306*T2 + 0.00001236*T3
	F := 21.2964 + 390.67050646*float64(k) - 0.0016528*T2 - 0.00000239*T3
	C1 := (0.1734-0.000393*T)*math.Sin(M*dr) + 0.0021*math.Sin(2*dr*M)
	C1 = C1 - 0.4068*math.Sin(Mpr*dr) + 0.0161*math.Sin(dr*2*Mpr)
	C1 = C1 - 0.0004*math.Sin(dr*3*Mpr)
	C1 = C1 + 0.0104*math.Sin(dr*2*F) - 0.0051*math.Sin(dr*(M+Mpr))
	C1 = C1 - 0.0074*math.Sin(dr*(M-Mpr)) + 0.0004*math.Sin(dr*(2*F+M))
	C1 = C1 - 0.0004*math.Sin(dr*(2*F-M)) - 0.0006*math.Sin(dr*(2*F+Mpr))
	C1 = C1 + 0.0010*math.Sin(dr*(2*F-Mpr)) + 0.0005*math.Sin(dr*(2*Mpr+M))

	deltat := 0.0
	if T < -11 {
		deltat = 0.001 + 0.000839*T + 0.0002261*T2 - 0.00000845*T3 - 0.000000081*T*T3
	} else {
		deltat = -0.000278 + 0.000265*T + 0.000262*T2
	}
	JdNew := Jd1 + C1 - deltat
	return int(JdNew + 0.5 + timeZone/24.0)
}

// getSunLongitude calculates sun longitude at Julian day number
func (c *Converter) getSunLongitude(jdn int) int {
	T := (float64(jdn) - 2451545.5 - timeZone/24.0) / 36525.0
	T2 := T * T
	dr := math.Pi / 180.0
	M := 357.52910 + 35999.05030*T - 0.0001559*T2 - 0.00000048*T*T2
	L0 := 280.46645 + 36000.76983*T + 0.0003032*T2
	DL := (1.914600 - 0.004817*T - 0.000014*T2) * math.Sin(dr*M)
	DL = DL + (0.019993-0.000101*T)*math.Sin(dr*2*M) + 0.000290*math.Sin(dr*3*M)
	L := L0 + DL
	L = L - float64(int(L/360.0))*360.0
	return int(L / 30.0)
}

// getLunarMonth11 finds the day number of the 11th month of the lunar year
func (c *Converter) getLunarMonth11(yy int) int {
	off := c.jdFromDate(31, 12, yy) - 2415021
	k := int(float64(off) / 29.530588853)
	nm := c.getNewMoonDay(k)
	sunLong := c.getSunLongitude(nm)

	if sunLong >= 9 {
		nm = c.getNewMoonDay(k - 1)
	}
	return nm
}

// getLeapMonthOffset determines the leap month offset
func (c *Converter) getLeapMonthOffset(a11 int) int {
	k := int((float64(a11) - 2415021.076998695) / 29.530588853)
	last := 0
	i := 1
	arc := c.getSunLongitude(c.getNewMoonDay(k + i))

	for arc != last && i < 14 {
		last = arc
		i++
		arc = c.getSunLongitude(c.getNewMoonDay(k + i))
	}
	return i - 1
}
