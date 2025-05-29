package holidays

import "time"

type EasterSunday struct {
	year int
}

func NewEasterSunday(year int) *EasterSunday {
	return &EasterSunday{year: year}
}

func (e *EasterSunday) Name() string {
	return "Easter Sunday"
}

func (e *EasterSunday) Date() time.Time {
	return guassAgorithm(e.year)
}

func guassAgorithm(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
