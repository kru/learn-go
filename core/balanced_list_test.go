package core

import "testing"

type sum struct {
	input    []int32
	expected string
}

var sums = []sum{
	{[]int32{1, 1, 4, 1, 1}, "YES"},
}

func TestBalancedSums(t *testing.T) {
	for _, test := range sums {
		if got := BalancedSums(test.input); got != test.expected {
			t.Errorf("test input %v, got %s, test.expected %s", test.input, got, test.expected)
		}
	}
}
