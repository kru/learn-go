package core

import "testing"

type times struct {
	input    string
	expected string
}

var tms = []times{
	{"00:01:45AM", "00:01:45"},
	{"12:45:03PM", "12:45:03"},
	{"07:05:45PM", "19:05:45"},
	{"11:59:45PM", "23:59:45"},
	{"12:31:13AM", "00:31:13"},
	{"06:40:03AM", "06:40:03"},
}

func TestConvert(t *testing.T) {
	for _, time := range tms {
		if got := Convert(time.input); got != time.expected {
			t.Errorf("got %s, test.expected %s", got, time.expected)
		}
	}
}
