package core

import "testing"

type list struct {
	input    []int32
	ln       int32
	expected int32
}

var lists = []list{
	{
		[]int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7802, 3142, 9739, 5629, 5413, 7232},
		5,
		1335,
	},
}

func TestUnfairness(t *testing.T) {
	for _, test := range lists {
		if got := Unfairness(test.ln, test.input); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
