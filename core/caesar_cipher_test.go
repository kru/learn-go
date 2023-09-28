package core

import "testing"

type msg struct {
	input    string
	salt     int32
	expected string
}

var msgs = []msg{
	{"www.abc.xy", 87, "fff.jkl.gh"},
}

func TestCaesarCipher(t *testing.T) {
	for _, test := range msgs {
		if got := CaesarCipher(test.input, test.salt); got != test.expected {
			t.Fatalf("got %s, test expected %s", got, test.expected)
		}
	}
}
