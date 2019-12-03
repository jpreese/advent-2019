package main

import "testing"

func TestFuelRequirement(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{12}, expected: 2},
		{input: []int{14}, expected: 2},
		{input: []int{1969}, expected: 966},
		{input: []int{100756}, expected: 50346},
	}

	for _, tt := range testCases {
		actual, err := FuelRequirement(tt.input)
		if err != nil {
			t.Fatalf("calculating fuel requirements: %v", err)
		}

		if tt.expected != actual {
			t.Errorf("unexpected fuel sum. expected %v actual %v", tt.expected, actual)
		}
	}
}
