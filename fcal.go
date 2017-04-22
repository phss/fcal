package main

import (
	"flag"
	"github.com/phss/fcal/calendar"
	"log"
	"time"
)

func main() {
	date := parseOptions()
	c := calendar.CalendarMonthFrom(date)
	calendar.PrintMonth(c)
}

func parseOptions() time.Time {
	var dateStr string
	flag.StringVar(&dateStr, "d", time.Now().Format("2006-01-02"), "ISO date of the day or month to be displayed")
	flag.Parse()

	date, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		log.Fatal(err)
	}

	return date
}
