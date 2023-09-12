package core

import (
	"math"
)

func DiagonalDifference(matrixes [][]int32) int32 {
	var left, right int32

	for i, outer := range matrixes {
		for j, inner := range outer {
			if i == j {
				left = left + inner
			}
			if len(outer)-1-j == i {
				right = right + inner
			}
		}
	}
	return int32(math.Abs(float64(left) - float64(right)))
}
