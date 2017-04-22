package main

import (
	"flag"
	"github.com/phss/fcal/calendar"
	"log"
	"strings"
	"time"
)

func main() {
	date, highlightDay := parseOptions()
	c := calendar.CalendarMonthFrom(date)
	calendar.PrintMonth(c, highlightDay)
}

func parseOptions() (time.Time, bool) {
	format := "2006-01-02"
	highlightDay := true

	var dateStr string
	flag.StringVar(&dateStr, "d", time.Now().Format(format), "ISO date of the day or month to be displayed")
	flag.Parse()

	if strings.Count(dateStr, "-") == 1 {
		format = "2006-01"
		highlightDay = false
	}

	date, err := time.Parse(format, dateStr)

	if err != nil {
		log.Fatal(err)
	}

	return date, highlightDay
}
