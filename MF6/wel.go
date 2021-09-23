package MF6

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

func Wel(fileName string, data []fileData) error {
	fmt.Println("Inside Wel function")
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

	for i := 0; i < monthCount; i++ { // TODO: check to make sure this loops to last month, might need to be monthCount + 1
		// filter data to just the fDate
		filteredData, err := filterDataByDate(nextDate, data)
		if err != nil {
			// since it's reasonable for no data in a month, might be the blank stress period
		}

		_ = filteredData
		// TODO: write stressPeriod header
		// TODO: write the stressPeriod data
		// TODO: write stressPeriod footer

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
func stressPeriod(fDate time.Time, data []fileData) ([]string, error) {
	// TODO: finish this to write the data out using fileData methods

	return nil, nil
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
