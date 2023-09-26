package main

import (
	"fmt"

	"github.com/kru/learn-go/core"
)

func main() {

	arr := []int32{3, 5, 6, 0, 7}
	power := []int32{3, 1, 0, 2}
	max := core.MaximizePower(arr, power)
	fmt.Println(max)

	arr1 := []int32{2, 4, 2, 1, 6}
	power1 := []int32{4, 1, 1, 3}
	max1 := core.MaximizePower(arr1, power1)
	fmt.Println(max1)
}
