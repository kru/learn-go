package core

import "testing"

type brb struct {
	input    []int32
	expected int32
}

var brbs = []brb{
	{[]int32{2, 1, 5, 3, 4}, 3},
	{[]int32{2, 5, 1, 3, 4}, -1},
	{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, 7},
	{[]int32{1, 2, 5, 3, 7, 8, 6, 9, 4}, 8},
	{[]int32{5, 1, 2, 3, 7, 8, 6, 4}, -1},
	{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, 7},
	{[]int32{1, 2, 5, 3, 4, 7, 8, 6}, 4},
}

func TestMinimumBribe(t *testing.T) {
	for _, test := range brbs {
		t.Logf("input %v", test.input)
		if got := MinimumBribe(test.input); got != test.expected {
			t.Fatalf("test expected %d, got %d", test.expected, got)
		}
	}
}
