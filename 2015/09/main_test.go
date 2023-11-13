package main

import "testing"

var benchmark int

func TestPartOne(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect int
	}{
		{"1", "London to Dublin = 464\n" +
			"London to Belfast = 518\n" +
			"Dublin to Belfast = 141", 605},
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
		{"1", "London to Dublin = 464\n" +
			"London to Belfast = 518\n" +
			"Dublin to Belfast = 141", 982},
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
