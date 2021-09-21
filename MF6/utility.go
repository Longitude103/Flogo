package MF6

import (
	"errors"
	"time"
)

// This package is several utility functions used across the MF6 package.
type fileData interface {
	date() time.Time
	node() int
	value() float64
}

// firstLastDate receives a slice of structs with a date() method and returns the first and last date that is present in the data.
func firstLastDate(data []fileData) (firstDate time.Time, lastDate time.Time, err error) {
	fDate := time.Now()
	lDate := time.Now()
	initDate := fDate

	for _, d := range data {
		if fDate.After(d.date()) {
			fDate = d.date()
		}

		if lDate.Before(d.date()) {
			lDate = d.date()
		}
	}

	if initDate == fDate || initDate == lDate {
		return fDate, lDate, errors.New("didn't find a date of the data")
	}

	return fDate, lDate, nil
}

func monthsCountSince(startDate time.Time, endDate time.Time) int {
	months := 0
	month := startDate.Month()
	for startDate.Before(endDate) {
		startDate = startDate.Add(time.Hour * 24)
		nextMonth := startDate.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}
