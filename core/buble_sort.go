package core

import (
	"fmt"
	"time"
)

func less(v int64, n int64) bool {
	return n < v
}

func isSorted(arr []int64) bool {
	for j := 0; j < len(arr)-1; j++ {
		if less(arr[j], arr[j+1]) {
			return false
		}
	}
	return true
}

func BubleSort(arr []int64) {
	st := time.Now().Nanosecond()
	if isSorted(arr) {
		et := time.Now().Nanosecond()
		fmt.Printf("%s recursion took %d\n", "BubleSort", et-st)
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		if less(arr[i], arr[i+1]) {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	BubleSort(arr)
}

// r := rand.New(rand.NewSource(time.Now().UnixNano()))
// list := r.Perm(100 * 1000)
// var list64 []int64
// for i := 0; i < len(list); i++ {
// 	list64 = append(list64, int64(list[i]))
// }
