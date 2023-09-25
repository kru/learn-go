package main

import (
	"fmt"

	"github.com/kru/learn-go/core"
)

func main() {

	tree := core.New(8)
	ch := make(chan int32)
	go core.Walk(tree, ch)

	for val := range ch {
		fmt.Println(val)
	}

	t1 := core.New(8)
	t2 := core.New(8)

	same := core.Same(t1, t2)
	fmt.Println("t1 == t2", same)

}
