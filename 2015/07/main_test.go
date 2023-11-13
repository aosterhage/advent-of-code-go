package main

import (
	"testing"
)

var benchmark *Circuit

func TestPartOne(t *testing.T) {
	var subtests = []struct {
		name   string
		input  string
		expect *map[string]uint16
	}{
		{"1",
			"123 -> x\n" +
				"456 -> y\n" +
				"x AND y -> d\n" +
				"x OR y -> e\n" +
				"x LSHIFT 2 -> f\n" +
				"y RSHIFT 2 -> g\n" +
				"NOT x -> h\n" +
				"NOT y -> i",
			&map[string]uint16{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456}},
		{"2",
			"y AND 1 -> z\n" +
				"x -> y\n" +
				"3 -> x",
			&map[string]uint16{
				"x": 3,
				"y": 3,
				"z": 1}},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			c := NewCircuit()
			c.ParseInstructions(subtest.input)
			for k, v := range *subtest.expect {
				if v != c.GetSignal(k) {
					t.Errorf("key %v is %v but expected %v", k, c.GetSignal(k), v)
				}
			}
		})
	}
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmark := NewCircuit()
		benchmark.ParseInstructions(input)
	}
}
