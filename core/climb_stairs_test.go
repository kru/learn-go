package core

import "testing"

type stair struct {
	input    int64
	expected int64
}

var stairs = []stair{
	{3, 4},
	{4, 7},
	{5, 13},
}

func TestClimbsStairs(t *testing.T) {
	for _, test := range stairs {
		if got := ClimbStairs(test.input); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
