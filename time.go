package goutils

import "time"

// IsToday checks if given Time is Today
func TimeIsToday(t time.Time) bool {
	t1 := time.Now()
	return TimeSameDay(t1, t)
}

// TimeSameDay checks if given times has same day
func TimeSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

// TimeSameMonth checks if given times has same month
func TimeSameMonth(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// TimeSameYear checks if given times has same year
func TimeSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

func TimeBeginningOfDay(t time.Time, location string) time.Time {
	loc, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	beginningOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	return beginningOfDay
}

func TimeEndOfDay(t time.Time, location string) time.Time {
	loc, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	endOfDay := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, loc)
	return endOfDay
}

func TimeInLoc(t time.Time, location string) time.Time {
	loc, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	return t.In(loc)
}
