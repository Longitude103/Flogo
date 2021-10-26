package MF6

import (
	"fmt"
	"path/filepath"
)

// Rch is a function to create the RCH file and creates the final string for the file name. It calls the welRchCreator
// utility function to make the file.
func Rch(fileName string, data []FileData, path string, mDesc string) error {
	fn := fileName + ".RCH6"
	fullPath := filepath.Join(path, fn)

	if err := welRchCreator(false, fullPath, data, mDesc); err != nil {
		return err
	}

	return nil
}

// rchHeader is a function to write the MODFLOW 6 RCH header required for all files. This does have defaults in it.
// the MAXBOUND * is replaced later by the HeaderMod function.
func rchHeader(mDesc string, maxBound int) ([]string, error) {
	hd := []string{"# MODFLOW6 Recharge Package\n"}
	if mDesc != "" {
		mDesc = "# " + mDesc + "\n"
		hd = append(hd, mDesc)
	}

	options := []string{"BEGIN OPTIONS\n", "  AUXILIARY  RCHLAYER\n", "  SAVE_FLOWS\n", "END OPTIONS\n", "\n", "BEGIN DIMENSIONS\n"}

	hd = append(hd, options...)
	hd = append(hd, fmt.Sprintf("  MAXBOUND %d\n", maxBound), "END DIMENSIONS\n\n\n")

	return hd, nil
}
