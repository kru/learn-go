package core

import "fmt"

// Given a smaller strings S and a bigger string input
// Design an algorithm to find all permutation of the shorter string
// within the longer one.
func PermutationS(s, input string) int {
	var count int
	smap := make(map[rune]int)

	var bi int32
	for _, val := range s {
		smap[val] += 1
		if val > bi {
			bi = int32(val)
		}
	}

	k := len(s)
	fmt.Println(len(input))
	for i := 0; i < len(input) && i+k <= len(input); i++ {
		end := i + k
		if i == len(input)-len(s) {
			if isPerm(smap, input[i:end]) {
				count++
				return count
			}
		}
		if int32(input[end]) > bi {
			i = end
			continue
		}
		tmp := input[i:end]
		if isPerm(smap, tmp) {
			fmt.Println(tmp, i, string(input[i]))
			count++
		}
	}

	return count
}

func isPerm(s map[rune]int, inputPart string) bool {
	shown := make(map[rune]int)
	for i, val := range s {
		shown[i] = val
	}
	for _, val := range inputPart {
		if shown[val] == 0 {
			return false
		}
		if shown[val] > 0 {
			shown[val] -= 1
		}
	}

	for key := range shown {
		if shown[key] != 0 {
			return false
		}
	}

	return true
}
