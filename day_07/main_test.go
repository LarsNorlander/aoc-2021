package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []int
		expected int
	}{
		{
			name:     "Sample Case",
			inputs:   []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expected: 37,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := PartOne(test.inputs)
			if actual != test.expected {
				t.Logf("expected: %d but got: %d instead", test.expected, actual)
				t.FailNow()
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []int
		expected int
	}{
		{
			name:     "Sample Case",
			inputs:   []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expected: 168,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := PartTwo(test.inputs)
			if actual != test.expected {
				t.Logf("expected: %d but got: %d instead", test.expected, actual)
				t.FailNow()
			}
		})
	}
}
