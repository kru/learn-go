package main

import (
	"fmt"

	"github.com/kru/learn-go/core"
)

func main() {

	t := core.NewBTree(8)
	ch := make(chan int32)
	go func() {
		core.WalkBTree(t, ch)
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

	t2 := core.NewBTree(3)
	t3 := core.NewBTree(3)
	same := core.SameBTree(t2, t3)
	fmt.Println(same)
	// core.BestOutlets("seattle", 100)
}
