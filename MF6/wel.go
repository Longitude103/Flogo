package MF6

import (
	"fmt"
	"path/filepath"
)

// Wel is a function to create the WEL6 file and creates the final string for the file name. It calls the welRchCreator
// utility function to make the file.
func Wel(fileName string, data []FileData, path string, mDesc string) error {
	fn := fileName + ".WEL6"
	fullPath := filepath.Join(path, fn)

	if err := welRchCreator(true, fullPath, data, mDesc); err != nil {
		return err
	}

	return nil
}

// welHeader is a function to write the MODFLOW 6 WEL6 welHeader required for all files. This does have defaults in it.
// the MAXBOUND * is replaced later by the HeaderMod function.
func welHeader(mDesc string, maxBound int) ([]string, error) {
	hd := []string{"# MODFLOW6 Well Boundary Package\n"}
	if mDesc != "" {
		mDesc = "# " + mDesc + "\n"
		hd = append(hd, mDesc)
	}

	options := []string{"BEGIN OPTIONS\n", "  SAVE_FLOWS\n", "  AUTO_FLOW_REDUCE 1.000000e-01\n", "END OPTIONS\n", "\n", "BEGIN DIMENSIONS\n"}

	hd = append(hd, options...)
	hd = append(hd, fmt.Sprintf("  MAXBOUND %d\n", maxBound), "END DIMENSIONS\n\n\n")

	return hd, nil
}
