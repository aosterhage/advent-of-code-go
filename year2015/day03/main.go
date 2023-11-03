package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

type Coordinates struct {
	x int
	y int
}

func PartOne(input string) int {
	santaPosition := Coordinates{0, 0}
	houses := make(map[Coordinates]bool)
	houses[santaPosition] = true

	for _, r := range input {
		switch r {
		case '^':
			santaPosition.y += 1
		case 'v':
			santaPosition.y -= 1
		case '>':
			santaPosition.x += 1
		case '<':
			santaPosition.x -= 1
		}

		houses[santaPosition] = true
	}

	return len(houses)
}

func PartTwo(input string) int {
	santaPosition := Coordinates{0, 0}
	roboSantaPosition := Coordinates{0, 0}
	houses := make(map[Coordinates]bool)
	houses[santaPosition] = true

	for i, r := range input {
		var someSanta *Coordinates
		if i%2 == 0 {
			someSanta = &santaPosition
		} else {
			someSanta = &roboSantaPosition
		}

		switch r {
		case '^':
			someSanta.y += 1
		case 'v':
			someSanta.y -= 1
		case '>':
			someSanta.x += 1
		case '<':
			someSanta.x -= 1
		}

		houses[*someSanta] = true
	}

	return len(houses)
}
