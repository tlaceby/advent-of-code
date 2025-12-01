package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/tlaceby/aoc/helpers"
)

//go:embed input.txt
var input string

func main() {
	zeroCounts := 0
	dial := 50

	for _, command := range helpers.SplitNewLine(input) {
		var passedZeroCount = 0

		offset := computeMoveOffset(command)
		dial, passedZeroCount = computeMove(dial, offset)
		zeroCounts += passedZeroCount
	}

	fmt.Printf("\nCounted %d times the dial pointed at 0\n", zeroCounts)
}

func computeMove(dial int, offset int) (int, int) {
	passedZeroCount := 0
	currentPos := dial

	for i := 0; i < helpers.AbsInt(offset); i++ {
		if offset > 0 {
			currentPos++
			if currentPos > 99 {
				currentPos = 0
			}
		} else {
			currentPos--
			if currentPos < 0 {
				currentPos = 99
			}
		}

		if currentPos == 0 {
			passedZeroCount++
		}
	}

	return currentPos, passedZeroCount
}

func computeMoveOffset(command string) int {
	num, _ := strconv.Atoi(command[1:])
	if string(command[0]) == "L" {
		return -num
	}

	return num
}
