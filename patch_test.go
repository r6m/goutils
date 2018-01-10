package utils

import (
	"testing"
)

func TestMapExists(t *testing.T) {
	type args struct {
		m     map[string]interface{}
		items []string
	}

	items := map[string]interface{}{
		"item1": 1,
		"item2": 2,
		"item3": 3,
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"multiple items", args{items, []string{"item1", "item2"}}, true},
		{"one item", args{items, []string{"item3"}}, true},
		{"item not exists", args{items, []string{"item1", "item2", "item4"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapExists(tt.args.m, tt.args.items...); got != tt.want {
				t.Errorf("MapExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceExists(t *testing.T) {
	type args struct {
		slice interface{}
		item  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"string slice exist", args{[]string{"one", "two", "three"}, "one"}, true},
		{"string slice not exist", args{[]string{"one", "two", "three"}, "four"}, false},
		{"int slice exist", args{[]int{1, 2, 3}, 2}, true},
		{"int slice not exist", args{[]int{1, 2, 3}, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceExists(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("SliceExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
