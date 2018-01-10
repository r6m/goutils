package goutils

import (
	"reflect"
	"testing"
)

func TestRandInt(t *testing.T) {
	first := RandInt(10, 100)

	// test if returns int
	if reflect.TypeOf(first).Kind() != reflect.Int {
		t.Errorf("RandInt() = %v, should be int", first)
	}

	if !(first > 10 && first < 100) {
		t.Errorf("RandInt() should be between 10 and 100, but is %v", first)
	}

	if !(first > 10 && first < 100) {
		t.Errorf("RandInt() should be between 10 and 100, but is %v", first)
	}

	second := RandInt(10, 100)
	if second == first {
		t.Errorf("RandInt() = %v, should be random, but gives same value", second)
	}

}

func TestRandInt64(t *testing.T) {
	first := RandInt64()
	if reflect.TypeOf(first).Kind() != reflect.Int64 {
		t.Errorf("RandInt64() = %v, should be int64", first)
	}

	second := RandInt64()
	if second == first {
		t.Errorf("RandInt64() = %v, should be random, but gives same value", second)
	}
}

func TestRandString(t *testing.T) {
	first := RandString(32)
	if reflect.TypeOf(first).Kind() != reflect.String {
		t.Errorf("RandString() = %v, should be String", first)
	}

	if len(first) != 32 {
		t.Errorf("RandString(), length should be %d but is %d", 32, len(first))
	}

	second := RandString(32)
	if second == first {
		t.Errorf("RandString() = %v, should be random, but gives same value", second)
	}
}
