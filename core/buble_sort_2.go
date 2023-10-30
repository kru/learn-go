package core

import (
	"fmt"
	"time"
)

func BubleSort2[T NumberList](arr []T) {
	st := time.Now().Nanosecond()
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			fmt.Println("idx", j)
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	et := time.Now().Nanosecond()
	fmt.Printf("%s nested for loop took %d\n", "BubleSort2", et-st)
}
