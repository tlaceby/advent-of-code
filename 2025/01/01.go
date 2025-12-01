package main

import (
	"fmt"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var inputString string

func main() {
	zeroCounts := 0 // number of times after an instruction that the pointer is at zero
	cursor := 50

	fmt.Printf(" - The dial starts by pointing at 50.\n")
	for _, command := range strings.Split(inputString, "\n") {
		offset := computeMoveOffset(command)
		cursor = computeMove(cursor, offset)
		fmt.Printf(" - The dial is rotated %s to point at %d.\n", command, cursor)
		if cursor == 0 {
			zeroCounts++
		}
	}

	fmt.Printf("\nCounted %d positions where the dial is left at 0\n", zeroCounts)
	fmt.Printf("Actual Password = %d\n", cursor)
}

func computeMove(cursor int, offset int) int {
	newCursor := (cursor + offset) % 100
	if newCursor < 0 {
		newCursor += 100
	}

	return newCursor
}

func computeMoveOffset(command string) int {
	num, _ := strconv.Atoi(command[1:])
	if string(command[0]) == "L" {
		return -num
	}

	return num
}
