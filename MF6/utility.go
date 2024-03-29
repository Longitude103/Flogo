package MF6

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

// FileData This package is several utility functions used across the MF6 package.
type FileData interface {
	Date() time.Time
	Node() int
	Value() float64
	UseValue() bool
	RowCol() (int, int)
	ConvertToFtPDay() float64
	ConvertToFt3PDay() float64
}

// firstLastDate receives a slice of structs with a date() method and returns the first and last date that is present in the data.
func firstLastDate(data []FileData) (firstDate time.Time, lastDate time.Time, err error) {
	fDate := time.Now()
	var lDate time.Time
	initDate := fDate
	first := true

	for _, d := range data {
		if fDate.After(d.Date()) {
			fDate = d.Date()
			if first {
				lDate = fDate
				first = false
			}
		}

		if lDate.Before(d.Date()) {
			lDate = d.Date()
		}
	}

	if initDate == fDate || initDate == lDate {
		fmt.Println(fDate, lDate)
		return fDate, lDate, errors.New("didn't find a date of the data")
	}

	return fDate, lDate, nil
}

// monthsCountSince is a function that counts the number of months between two dates. The first month is counted and the
// last month is not, so 6/1 to 10/1 is 4 months in this function, it does not consider number of days in either start
// or end dates so 6/1 to 10/31 is still 4 months.
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

// spHeader creates the stress period welHeader for WEL6
func spHeader(period int) string {
	return fmt.Sprintf("BEGIN PERIOD %d\n", period)
}

// spFooter creates the stress period footer
func spFooter() string {
	return fmt.Sprint("END PERIOD\n\n")
}

// writeLines is a function to write lines of text to a writer from a slice of strings
func writeLines(writer *bufio.Writer, lines []string) error {
	for _, line := range lines {
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

// filterDataByDate is a function that filters the slice of FileData and then returns another slice of only those records
// that match the date passed into the function.
func filterDataByDate(dt time.Time, data []FileData) (rData []FileData, length int, dataPresent bool) {
	for _, d := range data {
		if d.Date() == dt {
			rData = append(rData, d)
		}
	}

	length = len(rData)

	if len(rData) == 0 {
		return rData, 0, false
	}
	return rData, length, true
}

// stressPeriod is a function to return a slice of strings that are the formatted stress period data
func stressPeriod(data []FileData, wel bool, Rc bool) (spData []string, err error) {
	if len(data) == 0 {
		return spData, errors.New("no data")
	}

	for _, d := range data {
		var s string
		if wel {
			if Rc {
				r, c := d.RowCol()
				if d.UseValue() {
					s = fmt.Sprintf(" 1 %d %d %e\n", r, c, d.Value())
				} else {
					s = fmt.Sprintf(" 1 %d %d %e\n", r, c, d.ConvertToFt3PDay())
				}

			} else {
				// wel file just write the node and value
				if d.UseValue() {
					s = fmt.Sprintf(" %d %e\n", d.Node(), d.Value())
				} else {
					s = fmt.Sprintf(" %d %e\n", d.Node(), d.ConvertToFt3PDay())
				}

			}
		} else {
			// rch file, need a layer number
			if Rc {
				r, c := d.RowCol()
				if d.UseValue() {
					s = fmt.Sprintf(" 1 %d %d %e  1\n", r, c, d.Value())
				} else {
					s = fmt.Sprintf(" 1 %d %d %e  1\n", r, c, d.ConvertToFtPDay())
				}
			} else {
				if d.UseValue() {
					s = fmt.Sprintf(" %d %e 1\n", d.Node(), d.Value()) // single layer only, can do future upgrade
				} else {
					s = fmt.Sprintf(" %d %e 1\n", d.Node(), d.ConvertToFtPDay()) // single layer only, can do future upgrade
				}
			}
		}

		spData = append(spData, s)
	}

	return spData, nil
}

// welRchCreator is a utility function to creat either the wel or rch files based on the node number method. This function
// calls the header creator based on the wel bool. It writes the file in the path given.
func welRchCreator(wel bool, fullFilePath string, data []FileData, mDesc string, modelStartDate time.Time, Rc bool) error {
	file, err := os.Create(fullFilePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	var fileLines []string

	// write the first period
	fDate, lDate, err := firstLastDate(data)
	if err != nil {
		return err
	}

	// Add model start date
	monthsSinceStart := 0
	if modelStartDate.Before(fDate) {
		// model started before first date of data, need to start period numbering from first data period number
		monthsSinceStart = monthsCountSince(modelStartDate, fDate)
	}

	monthCount := monthsCountSince(fDate, lDate) + monthsSinceStart
	nextDate := fDate
	maxBound := 0

	for i := 0; i < monthCount+1; i++ {
		_, length, _ := filterDataByDate(nextDate, data)
		if maxBound < length {
			maxBound = length
		}

		nextDate = nextDate.AddDate(0, 1, 0)
	}
	nextDate = fDate

	var hd []string
	var errHd error
	if wel {
		hd, errHd = welHeader(mDesc, maxBound)
	} else {
		hd, errHd = rchHeader(mDesc, maxBound)
	}
	if errHd != nil {
		return err
	}

	// writes the welHeader
	fileLines = append(fileLines, hd...)
	if err := writeLines(writer, fileLines); err != nil {
		return err
	}

	for i := monthsSinceStart; i < monthCount+1; i++ {
		var spData []string

		// filter data to just the fDate
		filteredData, _, dataPresent := filterDataByDate(nextDate, data)

		// stress period welHeader
		spData = append(spData, spHeader(i+1))
		if !dataPresent {
			// since it's reasonable for no data in a month, might be the blank stress period
			spData = append(spData, spFooter())
			if err2 := writeLines(writer, spData); err2 != nil {
				return err2
			}

			nextDate = nextDate.AddDate(0, 1, 0)
			continue
		}

		// stress period data
		d, err := stressPeriod(filteredData, wel, Rc)
		if err == nil {
			spData = append(spData, d...)
		}

		// stress period footer
		spData = append(spData, spFooter())
		if err1 := writeLines(writer, spData); err1 != nil {
			return err1
		}

		nextDate = nextDate.AddDate(0, 1, 0)
	}

	return nil
}
