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

func GridChallenge(grid []string) string {
	fmt.Println(grid)
	var sortedGrid []string
	for i := 0; i < len(grid); i++ {
		var isSorted bool
		var prev int32
		tmp := make(gri, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if prev > int32(grid[i][j]) {
				isSorted = false
			}
			tmp[j] = int32(grid[i][j])
			prev = int32(grid[i][j])
		}
		if !isSorted {
			sort.Sort(tmp)
		}
		sortedGrid = append(sortedGrid, string(tmp))
	}

	for k := 0; k < len(sortedGrid[0]); k++ {
		var prev int32
		for col := 0; col < len(sortedGrid); col++ {
			if prev > int32(sortedGrid[col][k]) {
				return "NO"
			}
			prev = int32(sortedGrid[col][k])

		}
	}

	return "YES"
}
