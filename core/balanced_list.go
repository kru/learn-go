package core

import "fmt"

func countRS(list []int32) int32 {
	var total int32
	for _, val := range list {
		total += val
	}
	return total
}

// Find the source of this DSA to add more tests
func BalancedSums(arr []int32) string {
	if len(arr) == 1 {
		return "YES"
	}
	var ls, rs int32
	rs = countRS(arr)
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(i, ls, rs, arr[i])
		if ls == rs {
			return "YES"
		}
		if ls == arr[i+1] {
			return "YES"
		}
		if ls == arr[i+1] {
			return "YES"
		}
		if ls == rs-arr[i] {
			return "YES"
		}
		ls += arr[i]
		rs = rs - arr[i]
	}
	return "NO"
}
