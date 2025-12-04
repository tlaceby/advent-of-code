package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var width int
var height int
var grid [][]byte

func main() {
	rows := strings.Split(input, "\n")
	height = len(rows)
	width = len(rows[0])

	for _, r := range rows {
		gridRow := []byte(r)
		grid = append(grid, gridRow)
	}

	removedCount := 0

	fmt.Printf("width: %d, height: %d\n\n", width, height)

	for {
		countBefore := removedCount
		for y := range height {
			for x := range height {
				adjacentRolls := getAdjacentCount(x, y)
				if adjacentRolls < 4 && grid[y][x] == '@' {
					removedCount += 1
					grid[y][x] = 'x'
				}
			}
		}

		if countBefore == removedCount {
			break
		}
	}

	fmt.Printf("Removed count: %d\n", removedCount)
}

func getAdjacentCount(x, y int) int {
	count := 0
	// top row
	if y-1 >= 0 {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= width {
				continue
			}

			char := grid[y-1][i]
			if char == '@' {
				count++
			}
		}
	}

	// left
	if x-1 >= 0 && grid[y][x-1] == '@' {
		count++
	}

	// right

	if x+1 < width && grid[y][x+1] == '@' {
		count++
	}

	// bottom row
	if y+1 < height {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= width {
				continue
			}

			char := grid[y+1][i]
			if char == '@' {
				count++
			}
		}
	}

	return count
}
