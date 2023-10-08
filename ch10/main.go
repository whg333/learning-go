package main

import (
	"fmt"
	"learning-go/whg/calendar"
	"learning-go/whg/check"
)

func main() {
	date1 := calendar.Date{}
	fmt.Println(date1)
	err := date1.SetYear(2023)
	check.CheckAndLog(err)
	err = date1.SetMonth(10)
	check.CheckAndLog(err)
	err = date1.SetDay(7)
	check.CheckAndLog(err)
	fmt.Println(date1)

	date2 := calendar.Date{}
	date2.Init(2022, 11, 28)
	fmt.Println(date2)
	err = date2.SetYear(2021)
	check.CheckAndLog(err)
	fmt.Println(date2)

	date3 := calendar.NewDate(2013, 9, 28)
	fmt.Println(date3)
	err = date3.SetDay(7)
	check.CheckAndLog(err)
	fmt.Println(date3)

	event := calendar.NewEvent("纪念日")
	event.Init(2023, 10, 8)
	fmt.Println(event)
	event.Date = date3
	fmt.Println(event)
}
