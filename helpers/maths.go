package helpers

import "math"

func AbsInt(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}

func FloorInt(num float64) int {
	return int(math.Floor(num))
}

func Divmod(x int, y int) (float64, int) {
	mod := x % y
	div := float64(x) / float64(y)

	return div, mod
}
