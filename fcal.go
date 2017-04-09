package main

import (
	"github.com/phss/fcal/calendar"
	"time"
)

func main() {
	c := calendar.CalendarMonthFrom(time.Now())
	calendar.PrintMonth(c)
}
