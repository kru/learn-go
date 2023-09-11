package core

import (
	"fmt"
	"math"
)

func TwoCrystalBall(crystals []int) int {
	floor := 0
	jump := int(math.Floor(math.Sqrt(float64(len(crystals)))))

	lastIdx := 0
	for i := 0; i < len(crystals); i = i + jump {
		if crystals[i] == 1 {
			lastIdx = i
			break
		}
	}

	for j := lastIdx - jump; j < len(crystals); j++ {
		if crystals[j] == 1 {
			fmt.Printf("Last crystal will be break at %d floor\n", j+1)
			floor = j + 1
			break
		}
	}
	return floor
}
