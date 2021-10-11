package MF6

func Rch(fileName string, data []FileData) error {
	fn := "./" + fileName + ".rch"

	if err := welRchCreator(false, fn, data); err != nil {
		return err
	}

	return nil
}

// header is a function to write the MODFLOW 6 WEL6 header required for all files. This does have defaults in it.
func rchHeader() ([]string, error) {
	hd := []string{"# MODFLOW6 Well Boundary Package\n"}
	options := []string{"BEGIN OPTIONS\n", "  AUXILIARY  RCHLAYER\n", "  SAVE_FLOWS\n", "END OPTIONS\n", "\n"}

	hd = append(hd, options...)
	dm := []string{"BEGIN DIMENSIONS\n", "  MAXBOUND 143508\n", "END DIMENSIONS\n\n\n"}

	hd = append(hd, dm...)
	return hd, nil
}
