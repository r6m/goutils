package utils

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
)

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func RandInt64() int64 {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

func RandString(n int) string {
	bytes := make([]byte, n/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
