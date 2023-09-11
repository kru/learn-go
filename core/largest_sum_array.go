package core

import "fmt"

func sum(list []int64) int64 {
	var total int64
	for i := 0; i < len(list); i++ {
		total = total + list[i]
	}
	return total
}

func LargestSumArr(data []int64) int64 {
	var maxSoFar, maxEndHere int64
	var s, start, end int

	for i := 0; i < len(data); i++ {
		maxEndHere += data[i]
		if maxSoFar < maxEndHere {
			maxSoFar = maxEndHere
			end = i
			start = s + 1
		}
		if maxEndHere < 0 {
			maxEndHere = 0
			s = i
		}
	}

	fmt.Printf("start index: %d, end index: %d\n", start, end)
	return maxSoFar

}
