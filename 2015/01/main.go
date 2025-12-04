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
	for _, ch := range strings.Split(input, "") {
		floor += helpers.Ternary(ch == "(", 1, -1)
	}

	fmt.Printf("Floor: %d\n", floor)
}
