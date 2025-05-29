package holidays

import (
	"testing"
	"time"
)

func TestFixedHoliday(t *testing.T) {
	holiday := NewFixedHoliday("Test Holiday", time.May, 15, 2025)

	if holiday.Name() != "Test Holiday" {
		t.Errorf("Expected holiday name 'Test Holiday', got %s", holiday.Name())
	}

	expected := time.Date(2025, time.May, 15, 0, 0, 0, 0, time.UTC)
	if !holiday.Date().Equal(expected) {
		t.Errorf("Expected date %v, got %v", expected, holiday.Date())
	}
}

func TestEasterSunday2025(t *testing.T) {
	easter := NewEasterSunday(2025).Date()
	expected := time.Date(2025, time.April, 20, 0, 0, 0, 0, time.UTC)

	if !easter.Equal(expected) {
		t.Errorf("Expected Easter Sunday 2025 to be %v, got %v", expected, easter)
	}
}

func TestMovableHoliday(t *testing.T) {
	goodFriday := NewMovableHoliday("Good Friday", 2025, -2).Date()
	expected := time.Date(2025, time.April, 18, 0, 0, 0, 0, time.UTC)

	if !goodFriday.Equal(expected) {
		t.Errorf("Expected Good Friday 2025 to be %v, got %v", expected, goodFriday)
	}
}

func TestCalendarHolidaysCount(t *testing.T) {
	cal := NewCalendar(2025)
	holidays := cal.GetHolidays()

	if len(holidays) < 10 {
		t.Errorf("Expected at least 10 holidays, got %d", len(holidays))
	}
}
