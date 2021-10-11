package Flogo

import (
	"github.com/Longitude103/Flogo/MF6"
	"testing"
	"time"
)

type testData struct {
	Dt       time.Time
	V        float64
	CellNode int
}

func (t testData) Date() time.Time {
	return t.Dt
}

func (t testData) Node() int {
	return t.CellNode
}

func (t testData) Value() float64 {
	return t.V
}

var d1 = testData{Dt: time.Now(), V: 100, CellNode: 1344}
var d2 = testData{Dt: time.Now(), V: 150, CellNode: 15674}
var d3 = testData{Dt: time.Now(), V: 200, CellNode: 4325}

var data = []MF6.FileData{d1, d2, d3}

// Testing package for the input file generator.
func TestInput(t *testing.T) {
	if err := Input(true, false, "testFile", data); err != nil {
		t.Error("Function produced an error")
	}

}
