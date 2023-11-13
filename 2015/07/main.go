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

func PartOne(input string) int {
	c := NewCircuit()
	c.ParseInstructions(input)
	return int(c.GetSignal("a"))
}

func PartTwo(input string) int {
	c := NewCircuit()
	c.ParseInstructions(input)
	c.signals["b"] = strconv.FormatUint(uint64(PartOne(input)), 10)
	return int(c.GetSignal("a"))
}

type Circuit struct {
	signals map[string]string
}

func NewCircuit() *Circuit {
	var c Circuit
	c.signals = make(map[string]string)
	return &c
}

func (c *Circuit) ParseInstructions(input string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		words := strings.Split(line, " ")
		wireIndex := slices.Index(words, "->")

		value := ""
		for i, word := range words {
			if i == wireIndex {
				break
			}

			if i != 0 {
				value += " "
			}
			value += word

		}
		c.signals[words[wireIndex+1]] = value
	}
}

func (c *Circuit) GetSignal(signal string) uint16 {
	num, err := strconv.ParseInt(signal, 10, 0)
	if err == nil {
		return uint16(num)
	}

	words := strings.Split(c.signals[signal], " ")
	var value uint16

	if len(words) == 1 {
		return c.GetSignal(words[0])
	} else if words[0] == "NOT" {
		value = ^c.GetSignal(words[1])
	} else if words[1] == "AND" {
		value = c.GetSignal(words[0]) & c.GetSignal(words[2])
	} else if words[1] == "OR" {
		value = c.GetSignal(words[0]) | c.GetSignal(words[2])
	} else if words[1] == "LSHIFT" {
		value = c.GetSignal(words[0]) << c.GetSignal(words[2])
	} else { // words[1] == "RSHIFT" {
		value = c.GetSignal(words[0]) >> c.GetSignal(words[2])
	}

	valueString := strconv.FormatUint(uint64(value), 10)
	c.signals[signal] = valueString
	return value
}
