package main

import "testing"

func TestLookAndSay(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect string
	}{
		{"1", "1", "11"},
		{"2", "11", "21"},
		{"3", "21", "1211"},
		{"4", "1211", "111221"},
		{"5", "111221", "312211"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			result := LookAndSay(&subtest.input)
			if *result != subtest.expect {
				t.Errorf("result is %v (%v) but expected %v", len(*result), *result, subtest.expect)
			}
		})
	}
}
