package core

import "testing"

type pair struct {
	input    []int32
	power    []int32
	expected int32
}

var pairs = []pair{
	{[]int32{3, 5, 6, 0, 7}, []int32{3, 1, 0, 2}, 25},
	{[]int32{2, 4, 2, 1, 6}, []int32{4, 1, 1, 3}, 20},
}

func TestMaximizePower(t *testing.T) {
	for _, test := range pairs {
		if got := MaximizePower(test.input, test.power); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
