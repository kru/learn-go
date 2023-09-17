package core

import "testing"

type query struct {
	a        []int32
	b        []int32
	k        int32
	expected string
}

var queries = []query{
	{[]int32{3, 1, 2}, []int32{7, 8, 9}, 10, "YES"},
	{[]int32{1, 2, 2, 1}, []int32{3, 3, 3, 4}, 5, "NO"},
	{[]int32{3, 1}, []int32{1, 3}, 4, "YES"},
	{[]int32{2, 3, 1, 1, 1}, []int32{1, 3, 4, 3, 3}, 5, "NO"},
	{[]int32{1, 5, 1, 4, 4, 2, 7, 1, 2, 2}, []int32{8, 7, 1, 7, 7, 4, 4, 3, 6, 7}, 9, "NO"},
	{[]int32{3, 6, 8, 5, 9, 9, 4, 8, 4, 7}, []int32{5, 1, 0, 1, 6, 4, 1, 7, 4, 3}, 9, "YES"},
	{[]int32{4, 4, 3, 2, 1, 4, 4, 3, 2, 4}, []int32{2, 3, 0, 1, 1, 3, 1, 0, 0, 2}, 4, "YES"},
}

func TestTwoArrays(t *testing.T) {
	for _, test := range queries {
		if got := TwoArrays(test.k, test.a, test.b); got != test.expected {
			t.Errorf("test input %v, got %s, test.expected %s", test.a, got, test.expected)
		}
	}
}
