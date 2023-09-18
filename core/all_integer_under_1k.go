package core

import (
	"math"
)

type PairFailed struct {
	c, d int
}

// FindAllIntUnder1k print all positive integer solution of a3 + b3 = c3 + d3, where a,b,c d 1 <= 1000
func FindAllIntUnder1k() map[int][]int {
	pairs := make(map[int][]int)
	max := 1000

	for c := 1; c <= max; c++ {
		for d := 1; d <= max; d++ {
			result := math.Pow(float64(c), 3) + math.Pow(float64(d), 3)
			var pairCD []int
			if result <= 1000 {
				pairCD = append(pairCD, c, d)
				pairs[int(result)] = pairCD
			}
		}
	}

	return pairs
}
