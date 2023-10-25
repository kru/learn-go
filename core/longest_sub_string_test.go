package core

import "testing"

type str struct {
	input    string
	expected int
}

var strs = []str{
	{"nndNfdfdfdf", 4},
	{"xrxdfghrraasdfghjklbb", 10},
}

func TestLongestSubString(t *testing.T) {
	for _, test := range strs {
		if got := LongestSubString(test.input); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
