package core

// Given an array arr of n non-negative integers and an array power of k integers where k is an even number, perform the following operation in steps.

// • Select two integers and j such that 0 <= j <= |power|. In each operation, i and j are selected independently.
// • if power[i] < power[j] add the summation of the subarray arr[power[i]..power[j]] the power.
// • if power[i] > power[j] add the summation of the subarray arr[power[j]..power[i]] to the power.
// • Delete i-th and j-th elements from power, and its length is reduced by 2.

// Starting at 0 initial power, maximize the final power after exactly k/2 operations. Return that maximum power modulo (10^9 + 7).

// Note: Subarray arr[l...r] denotes the subarray formed by the elements arr[i], arr[i+ 1], ... arr[r-1], arr[i].

// Example
// arr= [3, 5, 6, 0, 7] and power = [3, 1, 0, 2].
// k=4, the size of power, so perform k/2 = 2 operations.

// One of the optimal approach is shown.
// • Select i=0, j=2. Here, power[0] = 3 and power[2] = 0. Add the sum of subarray arr[0...3]. (3 + 5 + 6 + 0) = 14 to
// the power, 0 + 14 = 14, Remove the two elements from power, and power is [1, 2].

// • Select i=0, j=1. Here, power[0] = 1 and power[1] = 2. Add the sum of subarray arr[1..2], 5 + 6 = 11, so power is
// 14 + 11 = 25.

func filter(power []int32, i, j int) []int32 {
	var pair []int32
	for idx := range power {
		if idx != i && idx != j {
			pair = append(pair, power[idx])
		}
	}
	return pair
}

func counter(arr []int32, pi, pj int32) int32 {
	var total int32
	if pi > pj {
		pi, pj = pj, pi
	}

	res := arr[pi : pj+1]
	for _, val := range res {
		total = total + val
	}

	return total
}

func MaximizePower(arr []int32, power []int32) int32 {
	var max int32

	for i := 0; i < len(power)/2; i++ {
		for j := 1; j < len(power); j++ {
			if i == j {
				continue
			}
			p1 := counter(arr, power[i], power[j])
			pair := filter(power, i, j)
			p2 := counter(arr, pair[0], pair[1])
			total := p1 + p2
			if total > max {
				max = total
			}
		}
	}

	return max
}
