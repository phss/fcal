package calendar

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
)

var titleStyle *color.Color = color.New(color.FgYellow)
var weekDaysStyle *color.Color = color.New(color.FgBlue)
var weekendStyle *color.Color = color.New(color.FgCyan)
var normalStyle *color.Color = color.New(color.Reset)

func PrintMonth(c CalendarMonth) {
	header := fmt.Sprintf("%s %d", c.Month, c.Year)
	titleStyle.Printf("%*s\n", 10+len(header)/2, header)
	weekDaysStyle.Println("Su Mo Tu We Th Fr Sa")

	var buffer bytes.Buffer
	weekDay := 0

	for i := 0; i < int(c.WeekDayStart); i++ {
		buffer.WriteString("   ")
		weekDay++
	}

	for day := 1; day <= c.LastDay; day++ {
		format := normalStyle
		if weekDay == 0 || weekDay == 6 {
			format = weekendStyle
		}
		buffer.WriteString(format.Sprintf("%2d ", day))
		weekDay++

		if weekDay == 7 || day == c.LastDay {
			weekDay = 0
			fmt.Println(buffer.String())
			buffer.Reset()
		}
	}
}
