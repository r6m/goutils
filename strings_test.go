package goutils

import (
	"testing"
)

func TestTruncateString(t *testing.T) {
	longString := RandString(100)
	truncated := TruncateString(longString, 10)

	// check if len is 10
	if len(truncated) != 10 {
		t.Errorf("truncated len must be 10, bot got %d", len(truncated))
	}

	if truncated != longString[:len(truncated)] {
		t.Errorf("truncated from beginning must be same as original's beginning, but got %s != %s", truncated, longString[:len(truncated)])
	}

	shortString := RandString(20)
	truncated = TruncateString(shortString, 25)

	if truncated != shortString {
		t.Errorf("truncate greater than origin string should return original string, bot got %d", len(truncated))
	}
}
