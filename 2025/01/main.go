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

// 10 L200 -> -190
// 6616
func computeMove(dial int, turn int) (int, int) {
	d := helpers.Ternary(turn > 0, 100, -100)
	div, mod := turn/d, turn%d

	count := div
	spread := dial + mod

	if turn < 0 {
		if dial != 0 && spread <= 0 {
			count += 1
		}
	} else {
		if spread >= 100 {
			count += 1
		}
	}

	dial = (dial + turn) % 100
	count = helpers.Ternary(dial == 0, count+1, count)
	return dial, count
}

func computeMoveOffset(command string) int {
	num, _ := strconv.Atoi(command[1:])
	if string(command[0]) == "L" {
		return -num
	}

	return num
}
