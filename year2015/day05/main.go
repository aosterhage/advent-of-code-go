package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/aosterhage/advent-of-code-go/util"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func PartOne(input string) int {
	potentialStrings := strings.Split(input, "\n")
	return len(util.Filter(potentialStrings, IsPartOneNiceString))
}

func IsPartOneNiceString(s string) bool {
	HasThreeVowels := func(s string) bool {
		return strings.Count(s, "a")+strings.Count(s, "e")+strings.Count(s, "i")+strings.Count(s, "o")+strings.Count(s, "u") >= 3
	}

	HasLetterAppearTwiceInARow := func(s string) bool {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				return true
			}
		}

		return false
	}

	HasSpecialStrings := func(s string) bool {
		return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
	}

	return HasThreeVowels(s) && HasLetterAppearTwiceInARow(s) && !HasSpecialStrings(s)
}

func PartTwo(input string) int {
	potentialStrings := strings.Split(input, "\n")
	return len(util.Filter(potentialStrings, IsPartTwoNiceString))
}

func IsPartTwoNiceString(s string) bool {
	HasLetterPairTwice := func(s string) bool {
		for i := 0; i < len(s)-1; i++ {
			if strings.Count(s, s[i:i+2]) > 1 {
				return true
			}
		}

		return false
	}

	HasLetterRepeatWithOneBetween := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				return true
			}
		}

		return false
	}

	return HasLetterPairTwice(s) && HasLetterRepeatWithOneBetween(s)
}
