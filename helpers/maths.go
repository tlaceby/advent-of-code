package helpers

import (
	"cmp"
	"math"
)

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

func Min[T cmp.Ordered](values []T) T {
	lowest := values[0]
	for _, v := range values {
		if v < lowest {
			lowest = v
		}
	}

	return lowest
}

func Max[T cmp.Ordered](values []T) T {
	highest := values[0]
	for _, v := range values {
		if v > highest {
			highest = v
		}
	}

	return highest
}
