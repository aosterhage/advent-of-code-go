package main

import "testing"

var benchmark int

func TestPartOne(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect bool
	}{
		{"1", "ugknbfddgicrmopn", true},
		{"2", "aaa", true},
		{"3", "jchzalrnumimnmhp", false},
		{"4", "haegwjzuvuyypxyu", false},
		{"5", "dvszwmarrgswjxmb", false},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			result := IsPartOneNiceString(subtest.input)
			if result != subtest.expect {
				t.Errorf("result is %v but expected %v", result, subtest.expect)
			}
		})
	}
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmark = PartOne(input)
	}
}

func TestPartTwo(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect bool
	}{
		{"1", "qjhvhtzxzqqjkmpb", true},
		{"2", "xxyxx", true},
		{"3", "uurcxstgmygtbstg", false},
		{"4", "ieodomkazucvgmuy", false},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			result := IsPartTwoNiceString(subtest.input)
			if result != subtest.expect {
				t.Errorf("result is %v but expected %v", result, subtest.expect)
			}
		})
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmark = PartTwo(input)
	}
}
