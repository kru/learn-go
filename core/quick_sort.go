package core

import "fmt"

func qs(ints []int32, low, high int) {
	if low >= high {
		return
	}
	p := partition(ints, low, high)

	qs(ints, low, p-1)
	qs(ints, p+1, high)
}

func partition(ints []int32, low, high int) int {
	pivot := ints[high]

	idx := low - 1

	for i := low; i < high; i++ {
		// this will swap everything less than equal pivot to the left side
		if ints[i] <= pivot {
			idx++
			ints[i], ints[idx] = ints[idx], ints[i]
		}
	}

	// this case we don't find anything less that pivot or
	// anything less than pivot have been moving to the left side
	// then we need to move our right after that lesser value
	// or our pivot should move the left side
	// []int32{9, 12, 5, 6, 10, 4} => []int32{4, 12, 5, 6, 10, 9}
	idx++
	ints[high], ints[idx] = ints[idx], pivot

	return idx
}

func QuickSort() {
	ints := []int32{9, 12, 5, 6, 10, 4}

	qs(ints, 0, len(ints)-1)

	fmt.Println(ints)
}
