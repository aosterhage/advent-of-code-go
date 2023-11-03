package main

import "testing"

var benchmark int

func TestPartOne(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect int
	}{
		{"1", "2x3x4", 58},
		{"2", "1x1x10", 43},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			l, w, h := ParseLengthWidthHeight(subtest.input)
			result := PartOne(l, w, h)
			if result != subtest.expect {
				t.Errorf("result is %v but expected %v", result, subtest.expect)
			}
		})
	}
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmark = Total(PartOne, input)
	}
}

func TestPartTwo(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect int
	}{
		{"1", "2x3x4", 34},
		{"2", "1x1x10", 14},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			l, w, h := ParseLengthWidthHeight(subtest.input)
			result := PartTwo(l, w, h)
			if result != subtest.expect {
				t.Errorf("result is %v but expected %v", result, subtest.expect)
			}
		})
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmark = Total(PartTwo, input)
	}
}
