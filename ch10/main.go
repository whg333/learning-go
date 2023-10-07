package main

import (
	"fmt"
	"learning-go/whg/calendar"
	"learning-go/whg/check"
)

func main() {
	date := calendar.Date{}
	fmt.Println(date)
	err := date.SetYear(2023)
	check.CheckAndLog(err)
	err = date.SetMonth(10)
	check.CheckAndLog(err)
	err = date.SetDay(7)
	check.CheckAndLog(err)
	fmt.Println(date)
}
