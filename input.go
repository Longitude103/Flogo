package Flowgo

import "fmt"

// Input is a function that returns and error but otherwise creates a zip file with the MODFLOW WEL or RCH files inside it
// in the local directory.
func Input(WEL bool, RCH bool, data interface{}) error {
	if WEL {
		fmt.Println("Create a WEL File")
	}

	if RCH {
		fmt.Println("Create a RCH File")
	}

	return nil
}
