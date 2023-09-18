package core

func countBars(slices []int32) int32 {
	var count int32

	for i := 0; i < len(slices); i++ {
		count = count + slices[i]
	}
	return count
}

// s is length of chocholate bar
// d sum of s[i] based on m value
// m length of combination from s
func ChocholateBars(s []int32, d, m int32) int32 {
	var total int32

	if len(s) == 1 && d == s[0] {
		return 1
	}

	for i := 0; i < len(s)-1; i++ {
		var end int
		if i+int(m) > len(s) {
			end = len(s)
		} else {
			end = i + int(m)
		}
		count := countBars(s[i:end])
		if count == d {
			total++
		}
	}

	return total
}
