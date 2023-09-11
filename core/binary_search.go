package core

import "fmt"

func BinarySearch(list []int32, value int32) bool {
	lo := 0
	hi := len(list)

	for lo < hi {
		idx := lo + (hi-lo)/2
		v := list[idx]
		fmt.Printf("index: %d; val: %d\n", idx, v)

		if v == value {
			fmt.Printf("Found the value %d\n", v)
			return true
		} else if v > value {
			fmt.Println("Searching on lower index")
			hi = idx
		} else {
			fmt.Println("Searching on higher index")
			lo = idx + 1
		}
	}

	fmt.Printf("value %d not found in the list", value)
	return false
}
