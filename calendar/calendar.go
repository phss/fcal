package calendar

import (
	"time"
)

type CalendarMonth struct {
	Month        string
	Year         int
	WeekDayStart int
	LastDay      int
}

func CalendarFrom(date time.Time) CalendarMonth {
	start := int(startOfMonth(date).Weekday())
	return CalendarMonth{Year: date.Year(), Month: date.Month().String(), WeekDayStart: start, LastDay: endOfMonth(date).Day()}
}

func startOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, 1-t.Day())
}

func endOfMonth(t time.Time) time.Time {
	return startOfMonth(t).AddDate(0, 1, -1)
}
