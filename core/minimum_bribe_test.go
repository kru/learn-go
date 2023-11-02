package core

import "testing"

type brb struct {
	input    []int32
	expected int32
}

var brbs = []brb{
	{[]int32{1, 2, 5, 3, 7, 8, 6, 4}, 7},
}

func TestMinimumBribe(t *testing.T) {
	for _, test := range brbs {
		if got := MinimumBribes(test.input); got != test.expected {
			t.Fatalf("test expected %d, got %d", test.expected, got)
		}
	}
}
