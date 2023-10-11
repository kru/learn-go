package core

import (
	"fmt"
	"strings"
)

var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func isPrime(n int32) bool {
	if n <= 1 {
		return false
	}

	i := int32(2)
	for n > 2 {
		if n%i == 0 {
			return false
		}
		i++
		n--
	}
	return true
}

func factor(n int32) int32 {
	for i := int32(3); i < 1000; i++ {
		if isPrime(i) && n%i == 1 {
			return i
		}
	}
	return n
}

func plantStrBombs(grid []string) []string {
	for row := range grid {
		grid[row] = strings.ReplaceAll(grid[row], ".", "O")
	}
	return grid
}

func plantRuneBoms(grid [][]rune) [][]rune {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = 'O'
		}
	}
	return grid
}

// @TODO: Revisit for optimization O(n)
// This problem not hard, but tricky
func BomberMan(n int32, grid []string) []string {
	bigNum := n
	if n == 1 {
		return grid
	}

	if n%2 == 0 {
		return plantStrBombs(grid)
	}
	n = factor(n)
	fmt.Println("PRime", n, bigNum)
	gridI := make([][]rune, len(grid))

	for i := range gridI {
		gridI[i] = []rune(grid[i])
	}

	for i := 2; i <= int(n); i++ {
		if i%2 == 0 {
			gridI = plantRuneBoms(gridI)
			continue
		}
		for row := range gridI {
			for col := range gridI[row] {
				for _, dir := range direction {
					if string(grid[row][col]) == "." {
						continue
					}
					x, y := row+dir[0], col+dir[1]
					gridI[row][col] = '.'
					if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[row])-1 {
						continue
					}
					gridI[x][y] = '.'
				}
			}
		}
		for row := range gridI {
			s := ""
			for col := range gridI[row] {
				s += string(gridI[row][col])
			}
			grid[row] = s
		}
	}

	for row := range gridI {
		s := ""
		for col := range gridI[row] {
			s += string(gridI[row][col])
		}
		grid[row] = s
	}

	return grid
}
