package main

import (
	"fmt"

	"github.com/kru/learn-go/core"
)

func main() {
	S := []int32{2, 2, 1, 3, 2}
	res := core.ChocholateBars(S, 4, 2)

	fmt.Println(res)

	B := []int32{4}
	re := core.ChocholateBars(B, 4, 1)
	fmt.Println(re)

	C := []int32{2, 2, 2, 1, 3, 2, 2, 3, 3, 1, 4, 1, 3, 2, 2, 1, 2, 2, 4, 2, 2, 3, 5, 3, 4, 3, 2, 1, 4, 5, 4}
	r := core.ChocholateBars(C, 10, 4)

	fmt.Println(r)
}
