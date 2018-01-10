package utils

import (
	"testing"
)

func pass(i interface{}) interface{} {
	return i
}
func TestToInt(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"int", args{pass(435424)}, 435424, false},
		{"negative int", args{pass(-435424)}, -435424, false},
		{"int64", args{pass(int64(435424))}, 435424, false},
		{"string", args{pass("123456")}, 123456, false},
		{"empty string", args{pass("")}, 0, true},
		{"float", args{pass(123.123)}, 123, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{"int", args{pass(435424)}, 435424, false},
		{"int", args{pass(-435424)}, -435424, false},
		{"int64", args{pass(int64(435424))}, 435424, false},
		{"string", args{pass("123456")}, 123456, false},
		{"float", args{pass(123.123)}, 123, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt64(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	type args struct {
		f    float64
		with bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"float", args{1234.123456789, true}, "1234.123456789"},
		{"float", args{1234.123, true}, "1234.123"},
		{"float", args{1234.123456789, false}, "1234"},
		{"float", args{12341231231231.234, false}, "12341231231231"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(tt.args.f, tt.args.with); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"string true", args{pass("true")}, true, false},
		{"string false", args{pass("false")}, false, false},
		{"string unknown", args{pass("unknown")}, false, true},

		{"int 1", args{pass(1)}, true, false},
		{"int 0", args{pass(0)}, false, false},
		{"int -1", args{pass(-1)}, false, true},

		{"int64 1", args{pass(int64(1))}, true, false},
		{"int64 0", args{pass(int64(0))}, false, false},
		{"int64 -1", args{pass(int64(-1))}, false, true},

		{"float 1.0", args{pass(1.0)}, true, false},
		{"float 0.0", args{pass(0.0)}, false, false},
		{"float -1.5", args{pass(-1.5)}, false, true},

		{"bool true", args{pass(true)}, true, false},
		{"bool false", args{pass(false)}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
