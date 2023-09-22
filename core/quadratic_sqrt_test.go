package core

import "testing"

type root struct {
	input    []float64
	expected []float64
}

var roots = []root{
	{[]float64{2, 10, 8}, []float64{-1.000000, -4.000000}},
}

func TestFindRoots(t *testing.T) {
	for _, test := range roots {
		x1, x2 := FindRoots(test.input[0], test.input[1], test.input[2])

		if x1 != test.expected[0] || x2 != test.expected[1] {
			t.Errorf("got x1: %.2f; x2: %.2f, expected %v", x1, x2, test.expected)
		}
	}
}
