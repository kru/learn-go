package core

import (
	"strings"
)

func LongestSubString(str string) int {
	substr := ""

	var lastIdx int
	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1] && !strings.Contains(substr, string(str[i])) {
			if i-lastIdx == 1 {
				substr += string(str[i])
				lastIdx = i
			}
			if i-lastIdx > 1 && len(str[i:]) > len(substr) {
				substr = ""
				lastIdx = i - 1
				i = i - 1
			}
			// not a good condition, this is second last character
			if i+1 == len(str)-2 && i+1-lastIdx == 1 {
				substr += string(str[i])
				lastIdx = i
			}
		}
	}

	return len(substr)
}
