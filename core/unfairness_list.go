package core

import (
	"sort"
)

// You will be given a list of integers, arr, and a single integer k.
// You must create an array of length k from elements of arr such that its unfairness is minimized.
// Call that array arr'. Unfairness of an array is calculated as max(arr') - max(arr')
// max denotes the largest integer in arr', min denotes the smallest integer in arr'

type sorted []int32

func (s sorted) Len() int {
	return len(s)
}

func (s sorted) Less(i, j int) bool {
	return s[j] > s[i]
}

func (s sorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func countUf(pairs []int32) int32 {
	return pairs[len(pairs)-1] - pairs[0]
}

func Unfairness(k int32, arr []int32) int32 {
	var list sorted
	for _, ar := range arr {
		list = append(list, ar)
	}

	sort.Sort(list)

	var result int32
	for i := 0; i < len(list); i++ {
		if i+int(k) > len(list) {
			break
		}

		tmp := countUf(list[i : i+int(k)])
		if i == 0 {
			result = tmp
		}
		if tmp < result {
			result = tmp
		}
	}

	return result
}
