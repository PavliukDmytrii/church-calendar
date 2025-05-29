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
			NewFixedHoliday("Christmas", time.December, 25, year),
			NewFixedHoliday("Epiphany", time.January, 6, year),
			NewEasterSunday(year),
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
