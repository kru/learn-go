package core

import "testing"

type perm1 struct {
	s   string
	l   string
	res int
}

var perms1 = []perm1{
	{
		"abbc",
		"cbabadcbbabbcbabaabccbabc",
		7,
	},
}

func TestPermutationS(t *testing.T) {
	for _, test := range perms1 {
		if got := PermutationS(test.s, test.l); got != test.res {
			t.Fatalf("got %d, test expected %d", got, test.res)
		}
	}
}
