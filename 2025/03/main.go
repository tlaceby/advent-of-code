package main

import (
	_ "embed"
	"fmt"

	"github.com/tlaceby/aoc/helpers"
)

//go:embed input.txt
var input string

const MAX_DIGITS = 12

func main() {
	sum := 0
	banks := helpers.SplitNewLine(input)
	for _, line := range banks {
		jolt := computeJolt(line)
		fmt.Printf("Line: %s -> Jolt: %d\n", line, jolt)
		sum += jolt
	}

	fmt.Printf("Sum: %d\n", sum)
}

// pick the largest digit in order using a greedy aproach
func computeJolt(bank string) int {
	digits := []byte(bank)
	pairs := make([]byte, MAX_DIGITS)
	firstChoiceIndex := 0

	for pos := range MAX_DIGITS {
		// if pos = 0, then last choice MUST leave room for MAX_DIGITS for future selections
		// if pos = MAX_DIGITS - 1, then it will point to the index of the last digit in the digits[]
		lastChoiceIndex := len(digits) - MAX_DIGITS + pos
		bestIndex := firstChoiceIndex
		bestDigit := digits[firstChoiceIndex]

		for i := firstChoiceIndex; i <= lastChoiceIndex; i++ {
			if digits[i] > bestDigit {
				bestDigit = digits[i]
				bestIndex = i
			}
		}

		pairs[pos] = bestDigit
		firstChoiceIndex = bestIndex + 1
	}

	return helpers.Atoi(string(pairs))
}
