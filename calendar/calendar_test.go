package calendar

import (
	"reflect"
	"testing"
	"time"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
)

var testCases = []struct {
	date     time.Time
	expected CalendarMonth
}{
	{
		date: time.Date(2017, 4, 10, 23, 0, 0, 0, time.UTC),
		expected: CalendarMonth{
			Month:        time.April,
			Year:         2017,
			WeekDayStart: time.Saturday,
			LastDay:      30,
			MarkedDays:   []int{10}},
	},
	{
		date: time.Date(2016, 2, 1, 12, 12, 12, 12, time.UTC),
		expected: CalendarMonth{
			Month:        time.February,
			Year:         2016,
			WeekDayStart: time.Monday,
			LastDay:      29,
			MarkedDays:   []int{1}},
	},
	{
		date: time.Date(1999, 12, 31, 23, 59, 59, 59, time.UTC),
		expected: CalendarMonth{
			Month:        time.December,
			Year:         1999,
			WeekDayStart: time.Wednesday,
			LastDay:      31,
			MarkedDays:   []int{31}},
	},
}

func TestCalendarMonth(t *testing.T) {
	for _, testCase := range testCases {
		observed := CalendarMonthFrom(testCase.date)
		if !reflect.DeepEqual(observed, testCase.expected) {
			t.Fatalf("CalendarMonthFrom() = %v, want %v", observed, testCase.expected)
		}
	}
}

func TestCalendarSpec(t *testing.T) {
	Convey("Marked days", t, func() {
		c := CalendarMonth{
			Month:        time.April,
			Year:         2017,
			WeekDayStart: time.Sunday,
			LastDay:      30,
			MarkedDays:   []int{16}}

		Convey("10th of April should not be a marked day", func() {
			So(c.IsMarkedDay(10), should.BeFalse)
		})

		Convey("16th of April should be a marked day", func() {
			So(c.IsMarkedDay(16), should.BeTrue)
		})
	})
}

