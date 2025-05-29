package holidays

import (
	"sort"
	"time"
)

type Calendar struct {
	year     int
	holidays []Holiday
}

func NewCalendar(year int) *Calendar {
	return &Calendar{
		year: year,
		holidays: []Holiday{
			//Fixed
			NewFixedHoliday("Circumcision of Jesus", time.January, 1, year),
			NewFixedHoliday("Epiphany", time.January, 6, year),
			NewFixedHoliday("Annunciation", time.March, 25, year),
			NewFixedHoliday("Transfiguration", time.August, 6, year),
			NewFixedHoliday("Holy Cross Day", time.September, 14, year),
			NewFixedHoliday("All Saints' Day", time.November, 1, year),
			NewFixedHoliday("Immaculate Conception", time.December, 8, year),
			NewFixedHoliday("Christmas", time.December, 25, year),
			NewFixedHoliday("St. Stephen's Day", time.December, 26, year),
			//Movable
			NewEasterSunday(year),
			NewMovableHoliday("Ash Wednesday", year, -46),
			NewMovableHoliday("Palm Sunday", year, -7),
			NewMovableHoliday("Maundy Thursday", year, -3),
			NewMovableHoliday("Good Friday", year, -2),
			NewMovableHoliday("Ascension Day", year, 39),
			NewMovableHoliday("Pentecost", year, 49),
			NewMovableHoliday("Trinity Sunday", year, 56),
			NewMovableHoliday("Corpus Christi", year, 60),
		},
	}
}

func (c *Calendar) GetHolidays() []Holiday {
	holidaysCopy := make([]Holiday, len(c.holidays))
	copy(holidaysCopy, c.holidays)

	sort.Slice(holidaysCopy, func(i, j int) bool {
		return holidaysCopy[i].Date().Before(holidaysCopy[j].Date())
	})

	return holidaysCopy
}
