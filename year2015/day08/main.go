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
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		total += len(line)
		total -= GetUnquotedLen(line)
	}

	return total
}

func PartTwo(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		total += GetEncodedLen(line)
		total -= len(line)
	}

	return total
}

func GetUnquotedLen(s string) int {
	unquoted := s[1 : len(s)-1]

	length := 0
	for i := 0; i < len(unquoted); length++ {
		current := unquoted[i]
		if current == '\\' && unquoted[i+1] == 'x' {
			i += 4
		} else if current == '\\' {
			i += 2
		} else {
			i++
		}
	}

	return length
}

func GetEncodedLen(s string) int {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	return len(s) + 2
}
