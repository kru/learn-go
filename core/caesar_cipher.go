package core

// A - Z (65-90), a - z (97-122)
func CaesarCipher(s string, k int32) string {
	sb := ""
	if k > 25 {
		k = k%25 - (k / 25)
	}
	for _, str := range s {
		if str < 65 || str > 122 {
			sb += string(str)
			continue
		}
		if str > 90 && str < 97 {
			sb += string(str)
			continue
		}
		if str <= 90 && str+k > 90 {
			n := str + k - 90
			str = 65 + (n - 1)
		} else if str <= 122 && str+k > 122 {
			n := str + k - 122
			str = 97 + (n - 1)
		} else {
			str = str + k
		}
		sb += string(str)
	}
	return sb
}
