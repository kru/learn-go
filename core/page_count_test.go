package core

import "testing"

type page struct {
	input    []int32
	expected int32
}

var pages = []page{
	{[]int32{37455, 29835}, 3810},
	{[]int32{6, 2}, 1},
	{[]int32{5, 4}, 0},
	{[]int32{6, 5}, 1},
}

func TestPageCount(t *testing.T) {
	for _, test := range pages {
		if got := PageCount(test.input[0], test.input[1]); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
