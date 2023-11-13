package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func PartOne(input string) int {
	return strings.Count(input, "(") - strings.Count(input, ")")
}

func PartTwo(input string) int {
	floor := 0

	for i, r := range input {
		if r == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			return i + 1
		}
	}

	return -1
}
