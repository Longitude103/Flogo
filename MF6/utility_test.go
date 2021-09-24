package MF6

import (
	"fmt"
	"testing"
	"time"
)

func Test_monthsCountSince(t *testing.T) {
	startDate := time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC)

	count := monthsCountSince(startDate, endDate)

	fmt.Println(count)

	if count != 4 {
		t.Error("count function not correct number of months")
	}
}

func Test_firstLastDate(t *testing.T) {
	f, l, e := firstLastDate(data)
	if e != nil {
		t.Errorf("Function returned an err of %s", e)
	}

	if f != time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC) {
		t.Error("First date is wrong")
	}

	if l != time.Date(2021, 6, 1, 0, 0, 0, 0, time.UTC) {
		t.Error("Last date is wrong")
	}
}
