package MF6

import (
	"github.com/Longitude103/Flogo/Utils"
	"testing"
	"time"
)

// tests of the wel file generator
type testRCHStruct struct {
	dt             time.Time
	nd             int
	recharge       float64
	rw             int
	clm            int
	convertedValue bool
	cellsize       float64
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

func (t testRCHStruct) RowCol() (int, int) {
	return t.rw, t.clm
}

func (t testRCHStruct) UseValue() bool {
	return t.convertedValue
}

func (t testRCHStruct) ConvertToFtPDay() float64 {
	return (t.recharge / t.cellsize) / float64(Utils.TimeExt{T: t.dt}.DaysInMonth())
}

func (t testRCHStruct) ConvertToFt3PDay() float64 {
	return (t.recharge * 43560) / float64(Utils.TimeExt{T: t.dt}.DaysInMonth()) * -1
}

var r0 = testRCHStruct{dt: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC), nd: 1,
	recharge: .000086978, convertedValue: true}
var r1 = testRCHStruct{dt: time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC), nd: 2,
	recharge: 0.01258775, convertedValue: true}
var r2 = testRCHStruct{dt: time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC), nd: 1,
	recharge: 1.235587441, convertedValue: true}
var r3 = testRCHStruct{dt: time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC), nd: 2,
	recharge: .00255864, convertedValue: true}
var r4 = testRCHStruct{dt: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC), nd: 3,
	recharge: 0.33658, convertedValue: true}
var r5 = testRCHStruct{dt: time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC), nd: 3,
	recharge: .000054884, convertedValue: true}

var rchTestData = []FileData{r0, r1, r2, r3, r4, r5}

func Test_Rch(t *testing.T) {
	err := Rch("testRch", rchTestData, ".", "test description", false)

	if err != nil {
		t.Errorf("Function errored with %s", err)
	}
}
