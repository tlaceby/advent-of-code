package helpers

import "strings"

func SplitNewLine(input string) []string {
	return strings.Split(input, "\n")
}
