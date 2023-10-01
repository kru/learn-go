package main

import (
	"fmt"

	"github.com/kru/learn-go/core"
)

func main() {

	res := core.SumXor(1000000000000000000)
	fmt.Println(res)
	// s := core.BalancedSums([]int32{1, 1, 4, 1, 1})
	// fmt.Println(s)
	// queries := [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}}
	// res := core.DynamicList(2, queries)
	// fmt.Printf("%v\n", res)

	// t := core.NewBTree(8)
	// ch := make(chan int32)
	// go func() {
	// 	core.WalkBTree(t, ch)
	// 	close(ch)
	// }()

	// for val := range ch {
	// 	fmt.Println(val)
	// }

	// t2 := core.NewBTree(3)
	// t3 := core.NewBTree(3)
	// same := core.SameBTree(t2, t3)
	// fmt.Println(same)
	// core.BestOutlets("seattle", 100)
}
