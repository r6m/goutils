package utils

import (
	"errors"
	"strconv"
	"strings"
)

// ToInt convert interface to int
func ToInt(i interface{}) (int, error) {
	if i == nil {
		return 0, errors.New("input is nil")
	}
	v := i
	if vint, ok := v.(int); ok {
		return vint, nil
	} else if vstring, sok := v.(string); sok {
		if vstring == "" {
			return 0, errors.New("could not convert string to int")
		}
		return strconv.Atoi(vstring)
	} else if vfloat, fok := v.(float64); fok {

		return strconv.Atoi(FloatToString(vfloat, false))
	} else if vint64, ok := v.(int64); ok {
		return int(vint64), nil
	}

	return 0, errors.New("could not convert to int")
}

func ToInt64(i interface{}) (int64, error) {
	n, err := ToInt(i)
	if err != nil {
		return 0, err
	}
	return int64(n), nil
}

func FloatToString(f float64, with bool) string {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	if !with {
		return strings.Split(str, ".")[0]
	}
	return str
}

func ToBool(i interface{}) (bool, error) {
	if i == nil {
		return false, errors.New("input is nil")
	}
	v := i
	if vbool, ok := v.(bool); ok {
		return vbool, nil
	} else if vint, ok := v.(int); ok {
		switch vint {
		case 0:
			return false, nil
		case 1:
			return true, nil
		}
	} else if vstring, sok := v.(string); sok {
		if vstring == "" {
			return false, errors.New("could not convert string to bool")
		}

		return strconv.ParseBool(vstring)
	} else if vfloat, fok := v.(float64); fok {
		return strconv.ParseBool(FloatToString(vfloat, false))
	} else if vint64, ok := v.(int64); ok {
		switch vint64 {
		case 0:
			return false, nil
		case 1:
			return true, nil
		}
	}

	return false, errors.New("could not convert to bool")
}
