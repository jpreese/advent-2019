package main

import "testing"

func TestOutput(t *testing.T) {
	testCases := []struct {
		input    []int
		position int
		expected int
	}{
		{input: []int{1, 0, 0, 0, 99}, position: 0, expected: 2},
		{input: []int{2, 3, 0, 3, 99}, position: 3, expected: 6},
		{input: []int{2, 4, 4, 5, 99, 0}, position: 5, expected: 9801},
		{input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, position: 0, expected: 30},
	}

	for _, tt := range testCases {
		actual := Output(tt.input)

		if tt.expected != actual[tt.position] {
			t.Errorf("unexpected value at position %v. expected %v actual %v", tt.position, tt.expected, actual)
		}
	}
}
