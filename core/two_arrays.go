package core

import (
	"sort"
)

type foo []int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func lessThan(k int32, ab int32) bool {
	return ab < k
}

func TwoArrays(k int32, A []int32, B []int32) string {
	var ai, bi foo
	for _, v := range A {
		ai = append(ai, v)
	}
	for _, v := range B {
		bi = append(bi, v)
	}
	sort.Sort(ai)
	sort.Sort(bi)

	cond := "YES"

	for i := 0; i < len(ai); i++ {
		if lessThan(k, ai[i]+bi[i]) {
			cond = "NO"
			break
		}
	}

	var count int

	for j := len(ai) - 1; j >= 0; j-- {
		if lessThan(k, ai[j]+bi[len(ai)-1-j]) {
			cond = "NO"
			break
		}
		if count == len(ai)-1 {
			cond = "YES"
			break
		}
		count++
	}

	return cond
}
