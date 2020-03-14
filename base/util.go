package base

import "strconv"

func ToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func ToFloat(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}
