package utils

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
