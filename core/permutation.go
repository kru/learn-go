package core

func Perm(str []rune, f func([]rune)) {
	perm(str, f, 0)
}

func perm(str []rune, f func([]rune), i int) {
	// fmt.Println(i)
	if i > len(str) {
		f(str)
		return
	}
	perm(str, f, i+1)
	for j := i + 1; j < len(str); j++ {
		str[i], str[j] = str[j], str[i]
		perm(str, f, i+1)
		str[i], str[j] = str[j], str[i]
	}
}
