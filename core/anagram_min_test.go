package core

import "testing"

type angr struct {
	input    string
	expected int32
}

var angrs = []angr{
	{"aaabbb", 3},
	{"ab", 1},
	{"abc", -1},
	{"mnop", 2},
	{"xyyx", 0},
	{"xaxbbbxx", 1},
	{"hhpddlnnsjfoyxpciioigvjqzfbpllssuj", 10},
	{"asdfjoieufoa", 3},
	{"fdhlvosfpafhalll", 5},
	{"mvdalvkiopaufl", 5},
	{"xulkowreuowzxgnhmiqekxhzistdocbnyozmnqthhpievvlj", 13},
}

func TestAnagram(t *testing.T) {
	for _, test := range angrs {
		if got := Anagram(test.input); got != test.expected {
			t.Fatalf("got %d, test expected %d", got, test.expected)
		}
	}
}
