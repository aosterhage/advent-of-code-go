package main

import "testing"

var benchmark int

func TestPartOne(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect int
	}{
		{"1", "turn on 0,0 through 999,999", 1_000_000},
		{"2", "toggle 0,0 through 999,0", 1_000},
		{"3", "turn on 0,0 through 999,999\nturn off 499,499 through 500,500", 999_996},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			result := PartOne(subtest.input)
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
		expect int
	}{
		{"1", "turn on 0,0 through 0,0", 1},
		{"2", "toggle 0,0 through 999,999", 2_000_000},
		{"3", "turn off 0,0 through 999,999", 0},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			result := PartTwo(subtest.input)
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
