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

func PartOne(s string) int {
	return strings.Count(s, "(") - strings.Count(s, ")")
}

func PartTwo(s string) int {
	floor := 0

	for i, r := range s {
		if r == '(' {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			// the problem wants the character's position, which is the index + 1
			return i + 1
		}
	}

	return -1
}
