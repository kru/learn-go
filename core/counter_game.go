package core

import (
	"fmt"
	"math"
)

func recCounter(n int64, cnt int) int {
	if n == 1 {
		return cnt
	}
	cnt++

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

// https://www.hackerrank.com/challenges/one-month-preparation-kit-counter-game/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-two
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
