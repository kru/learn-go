package core

import "math"

func FindRoots(a, b, c float64) (float64, float64) {
	d := b*b - 4*a*c
	x1 := (-b + math.Sqrt(d)) / (2 * a)

	x2 := (-b - math.Sqrt(d)) / (2 * a)

	return x1, x2
}
