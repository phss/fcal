package main

import "fmt"
import "time"

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

func main() {
	now := time.Now()
	//now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(now.String())

	thisMonth := now.Month()

	fmt.Println(thisMonth.String())
	fmt.Println(now.Weekday().String())
	fmt.Println(now.Weekday())

	dates := datesBetween(startOfMonth(now), endOfMonth(now))
	for _, date := range dates {
		fmt.Println(date)
	}
}
