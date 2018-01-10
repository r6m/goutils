package utils

import "reflect"

func MapExists(m map[string]interface{}, items ...string) bool {
	if len(items) == 0 {
		return true
	}

	var count int
	for _, i := range items {
		if _, ok := m[i]; ok {
			count++
		}
	}

	if count == len(items) {
		return true
	}

	return false
}

func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
