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
		expected: CalendarMonth{month: "April", year: 2017, weekDayStart: 6, lastDay: 30},
	},
	{
		date:     time.Date(2016, 2, 1, 12, 12, 12, 12, time.UTC),
		expected: CalendarMonth{month: "February", year: 2016, weekDayStart: 1, lastDay: 29},
	},
}

func TestCalendarMonth(t *testing.T) {
	for _, testCase := range testCases {
		observed := CalendarFrom(testCase.date)
		if observed != testCase.expected {
			t.Fatalf("CalendarFrom() = %v, want %v", observed, testCase.expected)
		}
	}
}
