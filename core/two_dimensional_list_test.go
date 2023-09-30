package core

import (
	"testing"
)

type table struct {
	height   int
	width    int
	queries  []string
	expected [][]int
}

var tables = []table{
	{
		height:   3,
		width:    5,
		queries:  []string{"v 1 2", "x 2 2", "v 1 2", "> 2 1", "x 2 3", "> 2 1", "< 2 0"},
		expected: [][]int{{2, 2}, {-1, -1}, {2, 3}, {2, 4}, {-1, -1}},
	},
}

func TestTwoDimensionalList(t *testing.T) {
	for _, test := range tables {
		got := TwoDimensionalList(test.height, test.width, test.queries)
		for i := range test.expected {
			for j := range test.expected[i] {
				if test.expected[i][j] != got[i][j] {
					t.Fatalf("got %v, test expected %v", got[i], test.expected[i])
				}
			}
		}
	}
}
