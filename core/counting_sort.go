package core

// example: input []int32{1,3,4,7,1,7}, based on value index {0,2,0,1,1,0,0,2} => {1,1,3,4,7,7}
func CountingSort(slices []int32) []int32 {

	res := make([]int32, 100) // we can modify 100 to length of slices

	for i := 0; i < len(slices); i++ {
		res[slices[i]] += 1
	}

	return res
}
