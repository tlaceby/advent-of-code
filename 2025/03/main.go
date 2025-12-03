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
	sum := 0
	banks := helpers.SplitNewLine(input)
	for _, line := range banks {
		h1, h2 := getHighestJolts(line)
		f := helpers.Aoti(h1 + h2)
		fmt.Printf("%s ---> (%s, %s) => %d\n", line, h1, h2, f)
		sum += f
	}

	fmt.Printf("Sum: %d\n", sum)
}

func getHighestJolts(bank string) (string, string) {
	voltages := strings.Split(bank, "")
	highestVoltagePairs := [2]string{"", ""}
	highestVoltage := -10000

	for i := range len(voltages) - 1 {
		v1 := voltages[i]
		for j := i + 1; j < len(voltages); j++ {
			v2 := voltages[j]

			totalVoltageStr := v1 + v2
			totalVoltage := helpers.Aoti(totalVoltageStr)
			if totalVoltage > highestVoltage {
				highestVoltage = totalVoltage
				highestVoltagePairs[0] = v1
				highestVoltagePairs[1] = v2
			}
		}
	}

	return highestVoltagePairs[0], highestVoltagePairs[1]
}
