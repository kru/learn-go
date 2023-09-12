package core

import (
	"fmt"
	"math"
)

func mapper(n int64) []int64 {
	var bits []int64
	res := make([]int64, 32)
	for n > 0 {
		r := n % 2

		n = (n - r) / 2
		bits = append(bits, r)
	}

	for j := 0; j < len(bits); j++ {
		res[len(res)-1-j] = bits[j]

	}
	return res
}

func fliper(bits []int64) []int64 {
	fliped := make([]int64, len(bits))

	for idx, bit := range bits {
		var flip int64
		if bit == 0 {
			flip = 1
		} else {
			flip = 0
		}
		fliped[idx] = flip
	}
	return fliped
}

func binarySum(arr []int64) int64 {
	var total int64
	for idx, val := range arr {
		total += int64(float64(val) * math.Pow(float64(2), float64(len(arr)-1-idx)))
	}
	return total
}

// Flip 32 bits of number representation
func FlippingBits(n int64) int64 {
	bits := mapper(n)
	fmt.Printf("mapper result: %v\n", bits)

	res := fliper(bits)
	fmt.Printf("fliper result: %v\n", res)

	fmt.Println()
	f := binarySum(res)
	fmt.Printf("summed flipped bits: %v\n", f)
	fmt.Println()
	return f
}
