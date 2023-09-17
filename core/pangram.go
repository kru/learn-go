package core

import (
	"strings"
)

func Pangram(s string) string {
	alphabet := "abcdefghijklmnopqrstupwxyz"

	for _, alp := range alphabet {
		if strings.Contains(strings.ToLower(s), string(alp)) {
			continue
		}
		return "not pangram"
	}

	return "pangram"
}
