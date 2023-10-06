package core

import (
	"strings"
)

func Anagram(s string) int32 {
	ln := len(s) / 2

	s1 := s[:ln]
	s2 := s[ln:]

	if s1 == s2 {
		return 0
	}
	if len(s1) != len(s2) {
		return -1
	}

	var notFound int32
	var sf string
	for i := 0; i < len(s1); i++ {
		if !strings.Contains(s2, string(s1[i])) {
			sf = sf + string(s1[i])
			notFound++
		} else {
			s2 = strings.Replace(s2, string(s1[i]), "", 1)
		}
	}

	return notFound
}
