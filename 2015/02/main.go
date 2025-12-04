package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tlaceby/aoc/helpers"
)

//go:embed input.txt
var input string

func main() {
	area := 0
	for _, d := range helpers.SplitNewLine(input) {
		split := strings.Split(d, "x")
		l := helpers.Atoi(split[0])
		w := helpers.Atoi(split[1])
		h := helpers.Atoi(split[2])

		sa, minSide := surfaceArea(l, w, h)
		area += sa + minSide
	}

	fmt.Printf("Area %d\n", area)
}

func surfaceArea(l, w, h int) (int, int) {
	s1 := l * w
	s2 := w * h
	s3 := h * l

	sa := (2 * s1) + (2 * s2) + (2 * s3)
	minSide := helpers.Min([]int{s1, s2, s3})
	return sa, minSide
}
