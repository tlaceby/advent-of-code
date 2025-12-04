package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var rows []string
var width int
var height int

func main() {
	count := 0
	rows = strings.Split(input, "\n")
	height = len(rows)
	width = len(rows[0])

	fmt.Printf("width: %d, height: %d\n\n", width, height)
	for y := range height {
		for x := range height {
			adjacentRolls := getAdjacentCount(x, y)
			if adjacentRolls < 4 && rows[y][x] == '@' {
				count++
			}
		}
	}

	fmt.Printf("Total count: %d\n", count)
}

func getAdjacentCount(x, y int) int {
	count := 0
	// top row
	if y-1 >= 0 {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= width {
				continue
			}

			char := rows[y-1][i]
			if char == '@' {
				count++
			}
		}
	}

	// left
	if x-1 >= 0 && rows[y][x-1] == '@' {
		count++
	}

	// right
	if x+1 < width && rows[y][x+1] == '@' {
		count++
	}

	// bottom row
	if y+1 < height {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= width {
				continue
			}

			char := rows[y+1][i]
			if char == '@' {
				count++
			}
		}
	}

	return count
}
