package core

import (
	"fmt"
	"math"
)

type xPair struct {
	x1, x2 float64
}

func quadSqrt(a, b, c float64, ch chan xPair) {
	d := b*b - 4*a*c
	x1 := (-b + math.Sqrt(d)) / (2 * a)

	x2 := (-b - math.Sqrt(d)) / (2 * a)

	ch <- xPair{x1: x1, x2: x2}
}

func Calculate(s1, s2 []float64) {
	ch := make(chan xPair)
	go quadSqrt(s1[0], s1[1], s1[2], ch)
	go quadSqrt(s2[0], s2[1], s2[2], ch)

	x, y := <-ch, <-ch

	fmt.Println(x.x1, x.x2)
	fmt.Println(y.x1, y.x2)
}
