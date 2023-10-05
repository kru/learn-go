package core

import (
	"fmt"
)

func getMax(a, b, c, d int32) int32 {
	var m int32
	if a > b {
		m = a
	} else {
		m = b
	}

	if m < c {
		m = c
	}

	if m < d {
		m = d
	}

	return m
}

func FlippingMatrix(matrix [][]int32) int32 {
	x := len(matrix) / 2
	var total int32
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Println(
				"matrix quadran 1: ", matrix[i][j], i, j,
				" quadran 2: ", matrix[2*x-1-i][j], i, j,
				" quadran 3: ", matrix[i][2*x-1-j], i, j,
				" quadran 4: ", matrix[2*x-1-i][2*x-1-j], i, j,
			)

			total += getMax(matrix[i][j], matrix[2*x-1-i][j], matrix[i][2*x-1-j], matrix[2*x-1-i][2*x-1-j])
		}
	}

	return total
}
