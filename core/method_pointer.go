package core

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v vertex) Abs(vs vertex) float64 {
	return math.Sqrt(v.X*vs.X + v.Y*vs.Y)
}

func (v vertex) Scale(f float64) vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Printf("v.x: %v, v.y: %v \n", v.X, v.Y)
	return v
}

func MethodPointer(x float64, y float64, s float64) float32 {
	v := vertex{x, y}
	vs := v.Scale(s)
	return float32(v.Abs(vs))
}
