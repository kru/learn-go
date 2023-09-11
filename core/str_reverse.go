package core

import (
	"strings"
)

func ReverseString(str string) string {
	var sb strings.Builder

	split := strings.Split(str, " ")
	for i := len(split) - 1; i >= 0; i-- {
		sb.WriteString(split[i] + " ")
	}

	return sb.String()
}
