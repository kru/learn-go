package core

import "testing"

type st struct {
	lowers   []int32
	highers  []int32
	expected int32
}

var sts = []st{
	{[]int32{2, 6}, []int32{24, 36}, 2},
	{[]int32{2, 4}, []int32{16, 32, 96}, 3},
	{[]int32{1}, []int32{100}, 9},
}

func TestBetweenTwoSets(t *testing.T) {
	for _, test := range sts {
		if got := BetweenTwoSets(test.lowers, test.highers); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
