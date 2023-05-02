package MF6

import (
	"fmt"
	"path/filepath"
	"regexp"
	"time"
)

// Wel is a function to create the WEL6 file and creates the final string for the file name. It calls the welRchCreator
// utility function to make the file.
func Wel(fileName string, data []FileData, path string, mDesc string, modelStartDate time.Time, Rc bool) error {
	fn := fileName + ".WEL6"
	fullPath := filepath.Join(path, fn)

	if err := welRchCreator(true, fullPath, data, mDesc, modelStartDate, Rc); err != nil {
		return err
	}

	return nil
}

// welHeader is a function to write the MODFLOW 6 WEL6 welHeader required for all files. This does have defaults in it.
// the MAXBOUND * is replaced later by the HeaderMod function.
func welHeader(mDesc string, maxBound int) ([]string, error) {
	hd := []string{"# MODFLOW6 Well Boundary Package\n"}
	if mDesc != "" {
		mDescComment := "# " + mDesc + "\n"
		hd = append(hd, mDescComment)
	}

	pattern := regexp.MustCompile(`\s+`)
	csvFileName := pattern.ReplaceAllString(mDesc, "_")
	optionCSV := fmt.Sprintf("  AUTO_FLOW_REDUCE_CSV FILEOUT %s.csv\n", csvFileName)

	options := []string{"BEGIN OPTIONS\n", "  SAVE_FLOWS\n", "  AUTO_FLOW_REDUCE 1.000000e-01\n", optionCSV, "END OPTIONS\n", "\n", "BEGIN DIMENSIONS\n"}

	hd = append(hd, options...)
	hd = append(hd, fmt.Sprintf("  MAXBOUND %d\n", maxBound), "END DIMENSIONS\n\n\n")

	return hd, nil
}
