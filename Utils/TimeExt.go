package Utils

import "time"

type TimeExt struct {
	T time.Time
	Y int
}

func (tm TimeExt) EndOfMonth() time.Time {
	y, m, _ := tm.T.Date()
	beginMonth := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)

	return beginMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (tm TimeExt) EndOfYear() time.Time {
	y, _, _ := tm.T.Date()
	beginYear := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC)

	return beginYear.AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (tm TimeExt) DaysInMonth() int {
	_, _, d := tm.EndOfMonth().Date()

	return d
}

func (tm TimeExt) DaysInYear() int {
	t := TimeExt{T: time.Date(tm.Y, 1, 1, 0, 0, 0, 0, time.UTC)}
	ey := t.EndOfYear()
	return ey.YearDay()
}
