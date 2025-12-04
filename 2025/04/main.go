package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {

	fmt.Printf("Input: %s\n", input)
}
