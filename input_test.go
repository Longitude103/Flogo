package Flogo

import (
	"testing"
	"time"
)

type testData struct {
	Dt       time.Time
	V        float64
	CellNode int
	rw       int
	clm      int
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

func (t testData) RowCol() (int, int) {
	return t.rw, t.clm
}

var d1 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 0.00008, CellNode: 1344}
var d2 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 0.00002, CellNode: 15674}
var d3 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 0.00003, CellNode: 4325}
var d4 = testData{Dt: time.Date(2020, time.Month(5), 1, 0, 0, 0, 0, time.UTC), V: 0.00012, CellNode: 4325}
var d5 = testData{Dt: time.Date(2020, time.Month(5), 1, 0, 0, 0, 0, time.UTC), V: 0.00054, CellNode: 4325}
var d6 = testData{Dt: time.Date(2020, time.Month(6), 1, 0, 0, 0, 0, time.UTC), V: 0.00125, CellNode: 4325}

var r1 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 0.00008, rw: 1, clm: 4}
var r2 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 0.00002, rw: 2, clm: 4}
var r3 = testData{Dt: time.Date(2020, time.Month(4), 1, 0, 0, 0, 0, time.UTC), V: 0.00003, rw: 1, clm: 5}
var r4 = testData{Dt: time.Date(2020, time.Month(5), 1, 0, 0, 0, 0, time.UTC), V: 0.00012, rw: 2, clm: 5}
var r5 = testData{Dt: time.Date(2020, time.Month(5), 1, 0, 0, 0, 0, time.UTC), V: 0.00054, rw: 3, clm: 4}
var r6 = testData{Dt: time.Date(2020, time.Month(6), 1, 0, 0, 0, 0, time.UTC), V: 0.00125, rw: 3, clm: 5}

var data []interface {
	Date() time.Time
	Node() int
	Value() float64
	RowCol() (int, int)
}

// Testing package for the input file generator.
func TestInput(t *testing.T) {
	data = append(data, d1, d2, d3, d4, d5, d6)

	if err := Input(true, false, false, "testFile", data, ".", "test description"); err != nil {
		t.Error("Function produced an error")
	}

}

func TestInput2(t *testing.T) {
	data = append(data, d1, d2, d3, d4, d5, d6)

	if err := Input(false, true, false, "testRCHFile", data, ".", ""); err != nil {
		t.Error("Function produced error with RCH")
	}
}

func TestRCHRowCol(t *testing.T) {
	data = append(data, r1, r2, r3, r4, r5, r6)

	if err := Input(false, true, true, "testRCHFile", data, ".", ""); err != nil {
		t.Error("Function produced error with RCH")
	}
}
