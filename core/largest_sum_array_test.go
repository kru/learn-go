package core

import "testing"

type slice struct {
	input    []int64
	expected int64
}

var sumList = []slice{
	{[]int64{-2, 6, -7, 8, -9, 5, 1, 17}, 23},
	{[]int64{-3, -1, 23, 1, -22, 8, -30, -17}, 24},
}

func TestLargestSumArr(t *testing.T) {
	for _, test := range sumList {
		if got := LargestSumArr(test.input); got != test.expected {
			t.Errorf("got %d, test.expected %d", got, test.expected)
		}
	}
}
