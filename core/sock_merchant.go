package core

// Find pair of element
// 10 20 20 10 10 30 50 10 20 => 3
// 1 1 3 1 2 1 3 3 3 3 => 4
func SockMerchant(n int32, ar []int32) int32 {
	var max, count int32
	for ai := 0; ai < len(ar); ai++ {
		if ar[ai] > max {
			max = ar[ai]
		}
	}
	sorted := make([]int32, max+1)
	for i := 0; i < len(ar); i++ {
		sorted[ar[i]] += 1
	}

	for k := 0; k < len(sorted); k++ {
		if sorted[k] == 0 {
			continue
		}
		count = count + (sorted[k] / 2)
	}

	return count
}
