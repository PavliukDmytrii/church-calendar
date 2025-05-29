package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/DmytriiPavliuk/church-calendar/holidays"
)

func main() {
	yearFlag := flag.Int("year", time.Now().Year(), "Year for the church calendar")
	flag.Parse()

	year := *yearFlag
	calendar := holidays.NewCalendar(year)

	fmt.Printf("Church Holidays for %d:\n", year)
	for _, h := range calendar.GetHolidays() {
		fmt.Printf("%-20s %s\n", h.Name(), h.Date().Format("Mon, Jan 2, 2006"))
	}
}
