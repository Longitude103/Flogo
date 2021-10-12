package Flogo

import (
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

var d1 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 100, CellNode: 1344}
var d2 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 150, CellNode: 15674}
var d3 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 200, CellNode: 4325}

var data []interface {
	Date() time.Time
	Node() int
	Value() float64
}

// Testing package for the input file generator.
func TestInput(t *testing.T) {
	data = append(data, d1, d2, d3)

	if err := Input(true, false, "testFile", data, "."); err != nil {
		t.Error("Function produced an error")
	}

}

func TestInput2(t *testing.T) {
	data = append(data, d1, d2, d3)

	if err := Input(false, true, "testRCHFile", data, "."); err != nil {
		t.Error("Function produced error with RCH")
	}
}
