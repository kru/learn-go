package core

import (
	"math"
	"strconv"
)

func SumXor(n int64) int64 {
	if n == 0 {
		return 1
	}

	var counter int64

	bins := strconv.FormatInt(n, 2)

	for _, bin := range bins {
		if string(bin) == "0" {
			counter++
		}
	}

	return int64(math.Pow(2, float64(counter)))
}
