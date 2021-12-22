package Flogo

import (
	"github.com/Longitude103/Flogo/MF6"
	"time"
)

// Input is a function that returns and error but otherwise creates a file of a MODFLOW WEL or RCH file
// in the local directory. Data must be passed as a slice of interface{} with the Date(), Node(), Value(), and RowCol() methods.
// Input arguments are as follows:
// -----------------------------------------------
// WEL => boolean if you want to make a wel file
// RCH => boolean if you want to make a RCH file
// Rc => boolean if you want to use row - column format (true) in list files or node (false)
// fileName => string file name of the file
// data => the data needs to be in a struct and include Data, Node, Value methods, but also rowColumn() method if Rc == true
// path => a string path of where to save the file
// mDesc => a string of the model description for the file header
func Input(WEL bool, RCH bool, Rc bool, fileName string, data []interface {
	Date() time.Time
	Node() int
	Value() float64
	RowCol() (int, int)
}, path string, mDesc string) error {
	dataInterface := make([]MF6.FileData, len(data))
	for i, v := range data {
		dataInterface[i] = v
	}

	if WEL {
		if err := MF6.Wel(fileName, dataInterface, path, mDesc, Rc); err != nil {
			return err
		}
	}

	if RCH {
		if err := MF6.Rch(fileName, dataInterface, path, mDesc, Rc); err != nil {
			return err
		}
	}

	return nil
}
