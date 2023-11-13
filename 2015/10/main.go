package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func PartOne(input string) int {
	answer := &input
	for i := 0; i < 40; i++ {
		answer = LookAndSay(answer)
	}

	return len(*answer)
}

func PartTwo(input string) int {
	answer := &input
	for i := 0; i < 50; i++ {
		answer = LookAndSay(answer)
	}

	return len(*answer)
}

func LookAndSay(s *string) *string {
	var num rune
	var count int64
	var answer strings.Builder

	for _, r := range *s {
		if count == 0 {
			num = r
			count++
		} else if r == num {
			count++
		} else {
			answer.WriteString(strconv.FormatInt(count, 10))
			answer.WriteRune(num)

			num = r
			count = 1
		}
	}
	answer.WriteString(strconv.FormatInt(count, 10))
	answer.WriteRune(num)

	answerString := answer.String()
	return &answerString
}
