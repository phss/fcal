package calendar

import (
	"time"
)

type CalendarMonth struct {
	Month        time.Month
	Year         int
	WeekDayStart time.Weekday
	LastDay      int
}

func CalendarMonthFrom(date time.Time) CalendarMonth {
	return CalendarMonth{
		Year:         date.Year(),
		Month:        date.Month(),
		WeekDayStart: startOfMonth(date).Weekday(),
		LastDay:      endOfMonth(date).Day()}
}

func startOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, 1-date.Day())
}

func endOfMonth(date time.Time) time.Time {
	return startOfMonth(date).AddDate(0, 1, -1)
}
