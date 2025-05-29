package holidays

import "time"

type MovableHoliday struct {
	name   string
	year   int
	offset int
}

func NewMovableHoliday(name string, year int, offset int) *MovableHoliday {
	return &MovableHoliday{name: name, year: year, offset: offset}
}

func (m *MovableHoliday) Name() string {
	return m.name
}

func (m *MovableHoliday) Date() time.Time {
	easter := guassAgorithm(m.year)
	return easter.AddDate(0, 0, m.offset)
}
