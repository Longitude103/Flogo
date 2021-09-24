package MF6

func Wel(fileName string, data []fileData) error {
	fn := "./" + fileName + ".WEL6"

	if err := welRchCreator(true, fn, data); err != nil {
		return err
	}

	return nil
}

// welHeader is a function to write the MODFLOW 6 WEL6 welHeader required for all files. This does have defaults in it.
func welHeader() ([]string, error) {
	hd := []string{"# MODFLOW6 Well Boundary Package\n"}
	options := []string{"BEGIN OPTIONS\n", "  SAVE_FLOWS\n", "  AUTO_FLOW_REDUCE 1.000000e-01\n", "END OPTIONS\n", "\n"}

	hd = append(hd, options...)
	dm := []string{"BEGIN DIMENSIONS\n", "  MAXBOUND 19995\n", "END DIMENSIONS\n\n\n"}

	hd = append(hd, dm...)
	return hd, nil
}
