package main

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (d *Date) Scan(src interface{}) error {
	dateStr := string(src.([]uint8))
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	d.Year = t.Year()
	d.Month = int(t.Month())
	d.Day = t.Day()
	return nil
}

func (d Date) Value() (driver.Value, error) {
	parse, err := time.Parse("2006-01-02", d.String())
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d Date) MarshalJSON() ([]byte, error) {
	result := fmt.Sprintf("\"%d-%02d-%02d\"", d.Year, d.Month, d.Day)
	return []byte(result), nil
}

func (d *Date) UnmarshalJSON(input []byte) error {
	inputStr := strings.Trim(string(input), "\"")
	dates := strings.Split(inputStr, "-")
	d.Year = parseInt(dates[0])
	d.Month = parseInt(dates[1])
	d.Day = parseInt(dates[2])
	return nil
}

func parseInt(i string) int {
	i2, err := strconv.Atoi(i)
	if err != nil {
		panic(err)
	}
	return i2
}
