package Flogo

import (
	"github.com/Longitude103/Flogo/MF6"
	"time"
)

// Input is a function that returns and error but otherwise creates a zip file with the MODFLOW WEL or RCH files inside it
// in the local directory. data must be passed as a slice of interface{} with the Date(), Node() and Value() methods.
func Input(WEL bool, RCH bool, fileName string, data []interface {
	Date() time.Time
	Node() int
	Value() float64
}, path string) error {
	dataInterface := make([]MF6.FileData, len(data))
	for i, v := range data {
		dataInterface[i] = v
	}

	if WEL {
		if err := MF6.Wel(fileName, dataInterface, path); err != nil {
			return err
		}
	}

	if RCH {
		if err := MF6.Rch(fileName, dataInterface, path); err != nil {
			return err
		}
	}

	return nil
}
