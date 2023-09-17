package main

import (
	"fmt"

	"github.com/kru/learn-go/core"
)

func main() {
	A := []int32{3, 1, 2}
	B := []int32{7, 8, 9}
	res := core.TwoArrays(10, A, B)

	fmt.Println(res)

	C := []int32{1, 2, 2, 1}
	D := []int32{3, 3, 3, 4}
	r := core.TwoArrays(5, C, D)

	fmt.Println(r)

	E := []int32{3, 6, 8, 5, 9, 9, 4, 8, 4, 7}
	F := []int32{5, 1, 0, 1, 6, 4, 1, 7, 4, 3}
	re := core.TwoArrays(9, E, F)

	fmt.Println(re)
}
