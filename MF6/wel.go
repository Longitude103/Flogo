package MF6

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

func Wel(fileName string, data []fileData) error {
	fn := "./" + fileName + ".WEL6"
	file, err := os.Create(fn)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	var fileLines []string

	hd, err := header()
	if err != nil {
		return err
	}

	// writes the header
	fileLines = append(fileLines, hd...)
	if err := writeLines(writer, fileLines); err != nil {
		return err
	}

	fileLines = nil

	// write the first period
	fDate, lDate, err := firstLastDate(data)
	if err != nil {
		return err
	}

	monthCount := monthsCountSince(fDate, lDate)
	nextDate := fDate

	for i := 0; i < monthCount+1; i++ {
		var spData []string

		// filter data to just the fDate
		filteredData, err := filterDataByDate(nextDate, data)
		if err != nil {
			// since it's reasonable for no data in a month, might be the blank stress period
		}

		// stress period header
		spData = append(spData, spHeader(i+1))

		// stress period data
		d, err := stressPeriod(filteredData)
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

// header is a function to write the MODFLOW 6 WEL6 header required for all files. This does have defaults in it.
func header() ([]string, error) {
	hd := []string{"# MODFLOW6 Well Boundary Package\n"}
	options := []string{"BEGIN OPTIONS\n", "  SAVE_FLOWS\n", "  AUTO_FLOW_REDUCE 1.000000e-01\n", "END OPTIONS\n", "\n"}

	hd = append(hd, options...)
	dm := []string{"BEGIN DIMENSIONS\n", "  MAXBOUND 19995\n", "END DIMENSIONS\n\n\n"}

	hd = append(hd, dm...)
	return hd, nil
}

// spHeader creates the stress period header for WEL6
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

// stressPeriod is a function to return a slice of strings that are the formatted stress period data
func stressPeriod(data []fileData) (spData []string, err error) {
	if len(data) == 0 {
		return spData, errors.New("no data")
	}

	for _, d := range data {
		s := fmt.Sprintf(" %d %f\n", d.node(), d.value())
		spData = append(spData, s)
	}

	return spData, nil
}

// filterDataByDate is a function that filters the slice of fileData and then returns another slice of only those records
// that match the date passed into the function.
func filterDataByDate(dt time.Time, data []fileData) (rData []fileData, err error) {
	for _, d := range data {
		if d.date() == dt {
			rData = append(rData, d)
		}
	}

	if len(rData) == 0 {
		return rData, errors.New("did not have data for that period")
	}

	return rData, nil
}
