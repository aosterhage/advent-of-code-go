package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

type Action int

const (
	On Action = iota
	Off
	Toggle
)

type Coordinates struct {
	x int
	y int
}

func ParseAction(ss *[]string) Action {
	word := (*ss)[0]
	slices.Delete(*ss, 0, 1)
	if strings.Contains(word, "toggle") {
		return Toggle
	}

	word = (*ss)[0]
	slices.Delete(*ss, 0, 1)
	if strings.Contains(word, "on") {
		return On
	}

	return Off
}

func ParseCoordinates(ss *[]string) Coordinates {
	word := (*ss)[0]
	slices.Delete(*ss, 0, 1)

	nums := strings.Split(word, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return Coordinates{x, y}
}

func ParseLine(line string) (Action, Coordinates, Coordinates) {
	words := strings.Split(line, " ")

	action := ParseAction(&words)
	rangeBegin := ParseCoordinates(&words)
	slices.Delete(words, 0, 1) // "through"
	rangeEnd := ParseCoordinates(&words)

	return action, rangeBegin, rangeEnd
}

func PartOne(input string) int {
	lights := [1_000][1_000]bool{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		action, rangeBegin, rangeEnd := ParseLine(line)
		for x := rangeBegin.x; x <= rangeEnd.x; x++ {
			for y := rangeBegin.y; y <= rangeEnd.y; y++ {
				if action == On {
					lights[x][y] = true
				} else if action == Off {
					lights[x][y] = false
				} else {
					lights[x][y] = !lights[x][y]
				}
			}
		}
	}

	numOn := 0
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			if lights[x][y] {
				numOn++
			}
		}
	}
	return numOn
}

func PartTwo(input string) int {
	lights := [1_000][1_000]int{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		action, rangeBegin, rangeEnd := ParseLine(line)
		for x := rangeBegin.x; x <= rangeEnd.x; x++ {
			for y := rangeBegin.y; y <= rangeEnd.y; y++ {
				if action == On {
					lights[x][y]++
				} else if action == Off {
					if lights[x][y] > 0 {
						lights[x][y]--
					}
				} else {
					lights[x][y] += 2
				}
			}
		}
	}

	totalBrightness := 0
	for x := 0; x < len(lights); x++ {
		for y := 0; y < len(lights[x]); y++ {
			totalBrightness += lights[x][y]
		}
	}
	return totalBrightness
}
