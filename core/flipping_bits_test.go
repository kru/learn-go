package core

import "testing"

type bit struct {
	input    int64
	expected int64
}

var bits = []bit{
	{0, 4294967295},
	{1, 4294967294},
	{5783, 4294961512},
}

func TestFlippingBits(t *testing.T) {
	for _, test := range bits {
		if got := FlippingBits(test.input); got != test.expected {
			t.Errorf("got %d, test.expected %d", got, test.expected)
		}
	}
}
