package main

import "testing"

var benchmark int

func TestGetAdventCoinNumber(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect int
	}{
		{"1", "abcdef", 609043},
		{"2", "pqrstuv", 1048970},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			result := GetAdventCoinNumber(subtest.input, 5)
			if result != subtest.expect {
				t.Errorf("result is %v but expected %v", result, subtest.expect)
			}
		})
	}
}

func BenchmarkGetAdventCoinNumber(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmark = GetAdventCoinNumber(input, 5)
	}
}
