package core

import "testing"

type gr struct {
	input    []string
	expected string
}

var grs = []gr{
	// {
	// 	[]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}, "YES",
	// },
	// {
	// 	[]string{"abc", "lmp", "qrt"}, "YES",
	// },
	{
		[]string{"mpxz", "abcd", "wlmf"}, "NO",
	},
	{
		[]string{"abc", "hjk", "mpq", "rtv"}, "YES",
	},
}

func TestGridChallenge(t *testing.T) {
	for _, test := range grs {
		if got := GridChallenge(test.input); got != test.expected {
			t.Fatalf("got %s, test expected %s", got, test.expected)
		}
	}
}
