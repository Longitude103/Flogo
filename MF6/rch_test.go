package MF6

import (
	"testing"
	"time"
)

// tests of the wel file generator
type testRCHStruct struct {
	dt       time.Time
	nd       int
	recharge float64
}

func (t testRCHStruct) Date() time.Time {
	return t.dt
}

func (t testRCHStruct) Node() int {
	return t.nd
}

func (t testRCHStruct) Value() float64 {
	return t.recharge
}

var r0 = testRCHStruct{dt: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC), nd: 1,
	recharge: .0045786}
var r1 = testRCHStruct{dt: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC), nd: 2,
	recharge: 0.01258775}
var r2 = testRCHStruct{dt: time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC), nd: 1,
	recharge: 1.235587441}
var r3 = testRCHStruct{dt: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC), nd: 2,
	recharge: .00255864}
var r4 = testRCHStruct{dt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC), nd: 3,
	recharge: 0.33658}
var r5 = testRCHStruct{dt: time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC), nd: 3,
	recharge: .000054884}

var rchTestData = []FileData{r0, r1, r2, r3, r4, r5}

func Test_Rch(t *testing.T) {
	err := Rch("testRch", rchTestData)

	if err != nil {
		t.Errorf("Function errored with %s", err)
	}
}
