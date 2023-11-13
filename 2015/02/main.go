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
	fmt.Printf("Part 1: %v\n", Total(PartOne, input))
	fmt.Printf("Part 2: %v\n", Total(PartTwo, input))
}

func PartOne(l int, w int, h int) int {
	sides := (2 * l * w) + (2 * w * h) + (2 * h * l)
	slack := min(l*w, w*h, h*l)
	return sides + slack
}

func PartTwo(l int, w int, h int) int {
	wrap := min(2*l+2*w, 2*w+2*h, 2*h+2*l)
	bow := l * w * h
	return wrap + bow
}

func Total(f func(int, int, int) int, input string) int {
	presents := strings.Split(input, "\n")

	total := 0
	for _, present := range presents {
		l, w, h := ParseLengthWidthHeight(present)
		total += f(l, w, h)
	}

	return total
}

func ParseLengthWidthHeight(s string) (int, int, int) {
	values := strings.Split(s, "x")

	l, _ := strconv.Atoi(values[0])
	w, _ := strconv.Atoi(values[1])
	h, _ := strconv.Atoi(values[2])

	return l, w, h
}
