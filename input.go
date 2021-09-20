package Flowgo

import (
	"fmt"
	"github.com/Longitude103/Flowgo/MF6"
)

// Input is a function that returns and error but otherwise creates a zip file with the MODFLOW WEL or RCH files inside it
// in the local directory.
func Input(WEL bool, RCH bool, fileName string, data interface{}) error {
	if WEL {
		fmt.Println("Create a WEL File")
		if err := MF6.Wel(fileName, data); err != nil {
			return err
		}
	}

	if RCH {
		fmt.Println("Create a RCH File")
		MF6.Rch(data)
	}

	return nil
}
