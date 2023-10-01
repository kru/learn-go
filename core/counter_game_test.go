package core

import "testing"

type cg struct {
	number   int64
	expected string
}

var cgs = []cg{
	{132, "Louise"},
	{6, "Richard"},
}

func TestCounterGame(t *testing.T) {
	for _, test := range cgs {
		if got := CounterGame(test.number); got != test.expected {
			t.Fatalf("got %s, test expected %s", got, test.expected)
		}
	}
}
