package calendar

import (
	"testing"
	"time"
)

var testCases = []struct {
	date     time.Time
	expected CalendarMonth
}{
	{
		date:     time.Date(2017, 4, 10, 23, 0, 0, 0, time.UTC),
		expected: CalendarMonth{Month: time.April, Year: 2017, WeekDayStart: time.Saturday, LastDay: 30},
	},
	{
		date:     time.Date(2016, 2, 1, 12, 12, 12, 12, time.UTC),
		expected: CalendarMonth{Month: time.February, Year: 2016, WeekDayStart: time.Monday, LastDay: 29},
	},
	{
		date:     time.Date(1999, 12, 31, 23, 59, 59, 59, time.UTC),
		expected: CalendarMonth{Month: time.December, Year: 1999, WeekDayStart: time.Wednesday, LastDay: 31},
	},
}

func TestCalendarMonth(t *testing.T) {
	for _, testCase := range testCases {
		observed := CalendarMonthFrom(testCase.date)
		if observed != testCase.expected {
			t.Fatalf("CalendarMonthFrom() = %v, want %v", observed, testCase.expected)
		}
	}
}
