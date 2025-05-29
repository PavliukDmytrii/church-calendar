package holidays

import "time"

type Calendar struct {
	year     int
	holidays []Holiday
}

func NewCalendar(year int) *Calendar {
	return &Calendar{
		year: year,
		holidays: []Holiday{
			NewFixedHoliday("Christmas", time.December, 25, year),
			NewFixedHoliday("Epiphany", time.January, 6, year),
			NewEasterSunday(year),
		},
	}
}

func (c *Calendar) GetHolidays() []Holiday {
	return c.holidays
}
