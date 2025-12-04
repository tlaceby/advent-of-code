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
	floor := 0
	for indx, ch := range strings.Split(input, "") {
		floor += helpers.Ternary(ch == "(", 1, -1)
		if floor == -1 {
			fmt.Printf("entered basement at %d\n", indx+1)
			return
		}
	}

	fmt.Printf("Floor: %d\n", floor)
}
