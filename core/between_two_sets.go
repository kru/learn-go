package core

import "fmt"

func BetweenTwoSets(lowers, highers []int32) int32 {
	var result []int32

	start := lowers[0]
	end := highers[len(highers)-1]

	for i := start; i <= end; i++ {
		onLow, onHigh := true, true
		for _, low := range lowers {
			if i%low != 0 {
				onLow = false
			}
		}
		for _, hi := range highers {
			if hi%i != 0 {
				onHigh = false
			}
		}

		if onLow && onHigh {
			result = append(result, i)
		}
	}
	fmt.Println(result)
	return int32(len(result))
}
