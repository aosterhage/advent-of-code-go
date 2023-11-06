package main

import "testing"

var benchmark int

func TestPartOne(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect int
	}{
		{"1", "\"\"", 2},
		{"2", "\"abc\"", 2},
		{"3", "\"aaa\\\"aaa\"", 3},
		{"4", "\"\\x27\"", 5},
		{"5", "\"a\\\"\"", 3},
		{"6", "\"a\\\\a\"", 3},
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
		{"1", "\"\"", 4},
		{"2", "\"abc\"", 4},
		{"3", "\"aaa\\\"aaa\"", 6},
		{"4", "\"\\x27\"", 5},
		{"5", "\"a\\\"\"", 6},
		{"6", "\"a\\\\a\"", 6},
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
