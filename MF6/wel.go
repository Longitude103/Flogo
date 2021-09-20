package MF6

import (
	"bufio"
	"fmt"
	"os"
)

type welData interface {
	printLine()
}

func Wel(fileName string, data ...interface{}) error {
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

	return nil
}

func header() ([]string, error) {
	hd := []string{"# MODFLOW6 Well Boundary Package\n"}
	options := []string{"BEGIN OPTIONS\n", "  SAVE_FLOWS\n", "  AUTO_FLOW_REDUCE 1.000000e-01\n", "END OPTIONS\n", "\n"}

	hd = append(hd, options...)
	dm := []string{"BEGIN DIMENSIONS\n", "  MAXBOUND 19995\n", "END DIMENSIONS\n\n\n"}

	hd = append(hd, dm...)
	return hd, nil
}

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
