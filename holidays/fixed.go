package holidays

import "time"

type FixedHoliday struct {
	name  string
	month time.Month
	day   int
	year  int
}

func NewFixedHoliday(name string, month time.Month, day, year int) *FixedHoliday {
	return &FixedHoliday{
		name:  name,
		month: month,
		day:   day,
		year:  year,
	}
}

func (f *FixedHoliday) Name() string {
	return f.name
}

func (f *FixedHoliday) Date() time.Time {
	return time.Date(f.year, f.month, f.day, 0, 0, 0, 0, time.UTC)
}
