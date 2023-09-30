package core

import (
	"fmt"
	"sort"
)

type gri []int32

func (g gri) Len() int {
	return len(g)
}

func (g gri) Less(i, j int) bool {
	return g[i] < g[j]
}

func (g gri) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func isContigous(list []int32) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}
	return true
}

func GridChallenge(grid []string) string {
	fmt.Println(grid)
	var strByte [][]int32
	columns := make([][]int32, len(grid[0]))
	for _, str := range grid {
		tmp := make(gri, len(str))
		for i, chr := range str {
			tmp[i] = chr
		}
		sort.Sort(tmp)

		strByte = append(strByte, tmp)

		for j := range tmp {
			columns[j] = append(columns[j], tmp[j])
		}

	}

	for i := 0; i < len(strByte); i++ {
		if !isContigous(strByte[i]) {
			return "NO"
		}
	}

	for j := 0; j < len(columns); j++ {
		if !isContigous(columns[j]) {
			return "NO"
		}
	}

	return "YES"
}
