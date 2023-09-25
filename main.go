package main

import (
	"github.com/kru/learn-go/core"
)

func main() {

	tree := core.New(8)
	ch := make(chan int32)
	core.Walk(tree, ch)

	// for val := range ch {
	// 	fmt.Println(val)
	// }

	// t1 := core.New(8)
	// t1.Traverse()
	// fmt.Println(s)
	// t2 := core.New(8)

	// same := core.Same(t1, t2)
	// fmt.Println("t1 == t2", same)

}
