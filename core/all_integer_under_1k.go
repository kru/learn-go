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
	//for c := 0; c <= max; c++ {
	//	for d := 0; d <= max; d++ {
	//		for a := 0; a <= max; a++ {
	//			for b := 0; b <= max; b++ {
	//				if math.Pow(float64(a), 3)+math.Pow(float64(b), 3) == math.Pow(float64(c), 3)+math.Pow(float64(d), 3) {
	//					fmt.Printf("%d + %d = %d + %d\n", a, b, c, d)
	//				}
	//			}
	//		}
	//	}
	//}

	//var pair []int32
	//result := int32(math.Pow(float64(c), 3)) + int32(math.Pow(float64(d), 3))
	//if result > 0 && int(result) <= max {
	//	pair = append(pair, int32(c), int32(d))
	//	pairs = append(pairs, pair)
	//}

	return pairs
}
