package main

import (
	"bytes"
	"fmt"
	"github.com/phss/fcal/calendar"
	"time"
)

func printCalendar(c calendar.CalendarMonth) {
	header := fmt.Sprintf("%s %d", c.Month, c.Year)
	fmt.Printf("%*s\n", 10+len(header)/2, header)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	var buffer bytes.Buffer
	weekDay := 0

	for i := 0; i < int(c.WeekDayStart); i++ {
		buffer.WriteString("   ")
		weekDay++
	}

	for day := 1; day <= c.LastDay; day++ {
		buffer.WriteString(fmt.Sprintf("%2d ", day))
		weekDay++

		if weekDay == 7 || day == c.LastDay {
			weekDay = 0
			fmt.Println(buffer.String())
			buffer.Reset()
		}
	}
}

func main() {
	c := calendar.CalendarFrom(time.Now())
	printCalendar(c)
}
