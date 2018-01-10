package utils

import (
	"testing"
	"time"
)

func TestTimeIsToday(t *testing.T) {
	if !TimeIsToday(time.Now()) {
		t.Error("should be today")
	}
}

func TestTimeSameDay(t *testing.T) {
	today := time.Now()
	if !TimeSameDay(today, time.Now()) {
		t.Error("should be today")
	}

	daysAgo := time.Now().Add(-5 * 24 * time.Hour)
	if TimeSameDay(daysAgo, time.Now()) {
		t.Error("should not be today")
	}
}

func TestTimeSameMonth(t *testing.T) {
	today := time.Now()
	if !TimeSameMonth(today, time.Now()) {
		t.Error("should be today")
	}
}

func TestTimeSameYear(t *testing.T) {
	today := time.Now()
	if !TimeSameYear(today, time.Now()) {
		t.Error("should be today")
	}
}
