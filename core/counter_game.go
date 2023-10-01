package core

import (
	"fmt"
	"math"
)

func recCounter(n int64, cnt int) int {
	cnt++
	if n == 1 {
		return cnt - 1
	}
	y := 0
	for int64(math.Pow(2, float64(y))) < n {
		y++
	}
	closeNum := int64(math.Pow(2, float64(y-1)))

	if closeNum == n {
		n = n / 2
	} else {
		n = n - closeNum
	}

	return recCounter(n, cnt)
}

func CounterGame(n int64) string {
	if n == 1 {
		return "Richard"
	}

	counter := recCounter(n, 0)
	fmt.Println(counter)

	if counter%2 == 0 {
		return "Richard"
	}

	return "Louise"
}
