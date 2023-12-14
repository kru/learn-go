package main

import (
	"github.com/kru/learn-go/core"
)

func main() {
	core.TrebuchetSpell()
	// graph := map[string][]string{
	// 	"a": {"c", "b"},
	// 	"b": {"d"},
	// 	"c": {"e"},
	// 	"d": {"f"},
	// 	"e": {""},
	// 	"f": {""},
	// }
	// core.AdjacencyList(graph)
	// list := []int32{1, 3, 2, 5, 4}
	// list := []int32{14, 5, 4, 2, 9, 8, 12}
	// core.ListMinLength(list, 6)

	// fmt.Println(queue)
	// core.MinimumBribes(queue)
	// core.InitDoublyLL()
	// core.SumXor(1000000000000000000)

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
