package convert

import (
	"strconv"
	"unicode"
)

func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToUint64(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
