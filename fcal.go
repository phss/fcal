package main

import (
	"bytes"
	"fmt"
	"github.com/phss/fcal/calendar"
	"time"
)

type cal struct {
	month string
	year  int
	weeks [][]int
}

func startOfMonth(t time.Time) time.Time {
	return t.AddDate(0, 0, 1-t.Day())
}

func endOfMonth(t time.Time) time.Time {
	return startOfMonth(t).AddDate(0, 1, -1)
}

func datesBetween(start, end time.Time) []time.Time {
	var dates []time.Time
	curr := start

	for curr.Before(end.AddDate(0, 0, 1)) {
		dates = append(dates, curr)
		curr = curr.AddDate(0, 0, 1)
	}

	return dates
}

func breakIntoWeeks(dates []time.Time) [][]time.Time {
	var weeks [][]time.Time
	currWeek := []time.Time{}

	for _, date := range dates {
		currWeek = append(currWeek, date)
		if int(date.Weekday()) == 6 {
			weeks = append(weeks, currWeek)
			currWeek = []time.Time{}
		}
	}
	weeks = append(weeks, currWeek)

	return weeks
}

func printCalendar(c cal) {
	header := fmt.Sprintf("%s %d", c.month, c.year)
	fmt.Printf("%*s\n", 10+len(header)/2, header)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	for j, week := range c.weeks {
		var buffer bytes.Buffer

		if j == 0 {
			for i := 0; i <= 6-len(week); i++ {
				buffer.WriteString("   ")
			}
		}

		for _, day := range week {
			buffer.WriteString(fmt.Sprintf("%2d ", day))
		}

		fmt.Println(buffer.String())
	}
}

func main() {
	fmt.Println(calendar.Thing())
	now := time.Now()
	header := fmt.Sprintf("%s %d", now.Month(), now.Year())
	fmt.Printf("%*s\n", 10+len(header)/2, header)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	dates := datesBetween(startOfMonth(now), endOfMonth(now))
	weeks := breakIntoWeeks(dates)
	for j, week := range weeks {
		var buffer bytes.Buffer

		if j == 0 {
			for i := 0; i <= 6-len(week); i++ {
				buffer.WriteString("   ")
			}
		}

		for _, date := range week {
			buffer.WriteString(fmt.Sprintf("%2d ", date.Day()))
		}

		fmt.Println(buffer.String())
	}

	c := cal{year: 2017,
		month: "April",
		weeks: [][]int{
			[]int{1},
			[]int{2, 3, 4, 5, 6, 7, 8},
			[]int{9, 10, 11, 12, 13, 14, 15},
			[]int{16, 17, 18, 19, 20, 21, 22},
			[]int{23, 24, 25, 26, 27, 28, 29},
			[]int{30},
		},
	}
	printCalendar(c)
}
