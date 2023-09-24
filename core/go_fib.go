package core

import (
	"fmt"
	"time"
)

func fib(n int, ch chan int64) {
	x, y := int64(0), int64(1)
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func Fibonacci() {
	ch := make(chan int64, 10)
	go fib(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}
}
