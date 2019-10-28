package goutils

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
)

// RandInt generates random int between min & max
func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandInt64 generates random int64
func RandInt64() int64 {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

// RandString generates random string with given length
func RandString(n int) string {
	bytes := make([]byte, n/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

const digitsBytes = "0123456789"
const lowerBytes = "abcdefghijklmnopqrstuvwxyz"
const upperBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type RandMode int

const (
	Digits = 1 << iota
	Lower
	Upper
	Humanize
)

// RandStringBytes returns a random string based on given flag.
// multiple flags can be combined: Digits|Lower
func RandStringBytes(n int, mode RandMode) string {
	source := make([]byte, 0)
	if mode&(Digits) != 0 {
		source = append(source, []byte(digitsBytes)...)
	}
	if mode&(Lower) != 0 {
		source = append(source, []byte(lowerBytes)...)
	}
	if mode&(Upper) != 0 {
		source = append(source, []byte(upperBytes)...)
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = source[rand.Intn(len(source))]
	}

	return string(b)
}
