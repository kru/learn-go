package core

func palin(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func PalindromeIndex(s string) int32 {
	if palin(s) {
		return -1
	}
	for idx := 0; idx <= len(s)-1; idx++ {
		sub := s[:idx] + s[idx+1:]
		if palin(sub) {
			return int32(idx)
		}
	}
	return -1
}
