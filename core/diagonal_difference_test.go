package core

import "testing"

type matrix struct {
	input    [][]int32
	expected int32
}

var matrixes = []matrix{
	{
		[][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}, 2,
	},
	{
		[][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, 15,
	},
	{
		[][]int32{{7, 2, 4}, {3, 1, 6}, {-10, 9, -20}}, 7,
	},
}

func TestDiagonalDifference(t *testing.T) {
	for _, test := range matrixes {
		if got := DiagonalDifference(test.input); got != test.expected {
			t.Errorf("got %d, test.expected %d", got, test.expected)
		}
	}
}
