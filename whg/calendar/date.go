package calendar

import (
	"errors"
	"learning-go/whg/check"
)

type Date struct {
	year  int
	month int
	day   int
}

func NewDate(year int, month int, day int) Date {
	date := Date{}
	date.Init(year, month, day)
	return date
}

func (d *Date) Init(year int, month int, day int) {
	check.CheckAndLog(d.SetYear(year))
	check.CheckAndLog(d.SetMonth(month))
	check.CheckAndLog(d.SetDay(day))
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) Year() int {
	return d.year
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
func (d *Date) Day() int {
	return d.day
}
