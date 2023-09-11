package core

import (
	"math"
)

type Abser interface {
	Abs() float64
}

func Iter(x float64, y float64) float64 {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VertexI{x, y}

	a = f
	a = &v

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser
	a = v
	return a.Abs()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexI struct {
	X, Y float64
}

func (v VertexI) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
