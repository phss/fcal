package calendar

import (
	"time"
)

type CalendarMonth struct {
	month        string
	year         int
	weekDayStart int
	lastDay      int
}

func CalendarFrom(date time.Time) CalendarMonth {
	start := int(startOfMonth(date).Weekday())
	return CalendarMonth{year: date.Year(), month: date.Month().String(), weekDayStart: start, lastDay: endOfMonth(date).Day()}
}

func startOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, 1-t.Day())
}

func endOfMonth(t time.Time) time.Time {
	return startOfMonth(t).AddDate(0, 1, -1)
}
