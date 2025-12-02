package helpers

import (
	"strconv"
	"strings"
)

func SplitNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func Aoti(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
