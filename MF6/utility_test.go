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
}
