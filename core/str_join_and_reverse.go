package core

import (
	"strings"
)

func JoinedAndReverse(strs ...string) (string, string) {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}

	joined := sb.String()
	sb.Reset()
	
	for i := len(strs) - 1; i >= 0; i-- {
		sb.WriteString(strs[i])
	}

	return joined, sb.String()
}
