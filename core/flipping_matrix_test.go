package core

import "testing"

type mtx struct {
	input    [][]int32
	expected int32
}

var mtxs = []mtx{
	{
		[][]int32{
			{112, 42, 83, 119},
			{56, 125, 56, 49},
			{15, 78, 101, 43},
			{62, 98, 114, 108},
		}, 414,
	}, {
		[][]int32{
			{107, 54, 128, 15},
			{12, 75, 110, 138},
			{100, 96, 34, 85},
			{75, 15, 28, 112},
		},
		488,
	},
}

func TestFlippingMatrix(t *testing.T) {
	for _, test := range mtxs {
		if got := FlippingMatrix(test.input); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
