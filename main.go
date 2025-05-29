package main

import (
	"fmt"
	"time"

	"github.com/DmytriiPavliuk/church-calendar/holidays"
)

func main() {
	year := time.Now().Year()
	calendar := holidays.NewCalendar(year)

	fmt.Printf("Church Holidays for %d:\n", year)
	for _, h := range calendar.GetHolidays() {
		fmt.Printf("%s: %s\n", h.Name(), h.Date().Format("2006-01-02"))
	}
}
