package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func PartOne(input string) int {
	return GetAdventCoinNumber(input, 5)
}

func PartTwo(input string) int {
	return GetAdventCoinNumber(input, 6)
}

func GetAdventCoinNumber(input string, desiredNumLeadingZeroes int) int {
	var leadingString string
	for i := 0; i < desiredNumLeadingZeroes; i++ {
		leadingString += "0"
	}

	num := 1
	for ; ; num++ {
		inputBytes := []byte(input + strconv.Itoa(num))
		hash := md5.Sum(inputBytes)
		hexString := hex.EncodeToString(hash[:])

		if hexString[:desiredNumLeadingZeroes] == leadingString {
			break
		}
	}

	return num
}
