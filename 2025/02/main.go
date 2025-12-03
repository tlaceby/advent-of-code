package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tlaceby/aoc/helpers"
)

//go:embed input.txt
var input string

type Range struct {
	Low  int
	High int
}

func main() {
	var sum = 0
	for _, rng := range createIDRanges() {
		for i := rng.Low; i <= rng.High; i++ {
			if isBadID(fmt.Sprint(i)) {
				sum += i
			}
		}
	}

	fmt.Printf("Sum of bad ID's: %d\n", sum)
}

func isBadID(id string) bool {
	for patternLen := 1; patternLen <= len(id)/2; patternLen++ {
		// constraints:
		//  - must be possible to achieve 2 repitions (or more)
		//  - pattern must repeat for all substrings in string
		// PLAN: iterate over all pattern lengths which satisfy the constraints above
		// for each length construct the pattern and check that all substrings of that length are of the pattern?
		//  If one fails, it's not a pattern, if all succeed, the ID is invalid

		if len(id)%patternLen != 0 {
			continue
		}

		pattern := id[:patternLen]
		isRepeating := true
		maxRepitions := len(id) / patternLen
		if maxRepitions < 2 {
			continue
		}

		for offset := patternLen; offset < len(id); offset += patternLen {
			if id[offset:offset+patternLen] != pattern {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}

func createIDRanges() []Range {
	ranges := make([]Range, 0)

	for _, idRange := range strings.Split(input, ",") {
		spread := strings.Split(idRange, "-")
		ranges = append(ranges, Range{
			Low:  helpers.Atoi(spread[0]),
			High: helpers.Atoi(spread[1]),
		})
	}

	return ranges
}
