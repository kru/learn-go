package core

import (
	"fmt"
	"strings"
)

func fillBombs(grid []string) []string {
	for row := range grid {
		grid[row] = strings.ReplaceAll(grid[row], ".", "O")
	}
	return grid
}

func plantBombs(grid [][]rune) [][]rune {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'O' {
				grid[i][j] = '*'
			} else {
				grid[i][j] = 'O'
			}
		}
	}
	return grid
}

// top {0, 1}
// right {1, 0}
// bottom {0, -1}
// left {-1, 0}
var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func detonate(grid [][]rune) [][]rune {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '*' {
				for _, dir := range direction {
					x, y := row+dir[0], col+dir[1]
					grid[row][col] = '.'
					if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[row])-1 {
						continue
					}
					if grid[x][y] == '*' {
						continue
					}
					grid[x][y] = '.'
				}
			}
		}
	}

	return grid
}

func convert(grid [][]rune) []string {
	var strs []string

	for row := 0; row < len(grid); row++ {
		str := ""
		for col := 0; col < len(grid[row]); col++ {
			str = str + string(grid[row][col])
		}
		strs = append(strs, str)
	}
	return strs
}

// @TODO: Revisit for optimization O(n)
// This problem not hard, but tricky
func BomberMan(n int32, grid []string) []string {
	if n == 1 {
		return grid
	}

	if n%2 == 0 {
		return fillBombs(grid)
	}

	gridI := make([][]rune, len(grid))

	for i := range gridI {
		gridI[i] = []rune(grid[i])
	}

	gridI = plantBombs(gridI)
	gridI = detonate(gridI)

	ans1 := convert(gridI)
	fmt.Println(ans1)
	if (n+1)%4 == 0 {
		return ans1
	}
	gridI = plantBombs(gridI)
	gridI = detonate(gridI)

	ans2 := convert(gridI)
	if (n+3)%4 == 0 {
		return ans2
	}
	return grid
}
