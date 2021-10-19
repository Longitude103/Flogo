package MF6

import "path/filepath"

func Rch(fileName string, data []FileData, path string, mDesc string) error {
	fn := fileName + ".rch"
	fullPath := filepath.Join(path, fn)

	if err := welRchCreator(false, fullPath, data, mDesc); err != nil {
		return err
	}

	return nil
}

// header is a function to write the MODFLOW 6 WEL6 header required for all files. This does have defaults in it.
func rchHeader(mDesc string) ([]string, error) {
	hd := []string{"# MODFLOW6 Recharge Package\n"}
	if mDesc != "" {
		mDesc = "# " + mDesc + "\n"
		hd = append(hd, mDesc)
	}

	options := []string{"BEGIN OPTIONS\n", "  AUXILIARY  RCHLAYER\n", "  SAVE_FLOWS\n", "END OPTIONS\n", "\n"}

	hd = append(hd, options...)
	dm := []string{"BEGIN DIMENSIONS\n", "  MAXBOUND 143508\n", "END DIMENSIONS\n\n\n"}

	hd = append(hd, dm...)
	return hd, nil
}
