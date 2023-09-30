package core

import (
	"fmt"
	"slices"
	"strconv"
)

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return int(num)
}

// Test from Code Signal
func TwoDimensionalList(h int, w int, queries []string) [][]int {
	grid := make([][]int, h)
	result := make([][]int, len(queries))
	for i := range grid {
		grid[i] = make([]int, w)
	}
	for idx, query := range queries {
		if string(query[0]) == "x" {
			row := toInt(string(query[2]))
			col := toInt(string(query[4]))
			if col > len(grid[0])-1 || row > len(grid)-1 {
				continue
			}
			grid[row][col] = 1
		}
		if string(query[0]) == "v" {
			row := toInt(string(query[2]))
			col := toInt(string(query[4]))
			if col > len(grid[0])-1 || row+1 > len(grid)-1 {
				result[idx] = []int{-1, -1}
				continue
			}

			for row < len(grid)-1 {
				row++
				if grid[col][row] == 1 {
					result[idx] = []int{-1, -1}
					continue
				}
				result[idx] = []int{row, col}
				break
			}
		}
		if string(query[0]) == "^" {
			row := toInt(string(query[2]))
			col := toInt(string(query[4]))
			if col < len(grid[0])-1 || row-1 < len(grid)-1 {
				result[idx] = []int{-1, -1}
				continue
			}
			for row > 0 {
				row--
				if grid[col][row] == 1 {
					result[idx] = []int{-1, -1}
					continue
				}
			}
			result[idx] = []int{row, col}
			break
		}
		if string(query[0]) == ">" {
			row := toInt(string(query[2]))
			col := toInt(string(query[4]))
			if col+1 > len(grid[0])-1 || row > len(grid)-1 {
				result[idx] = []int{-1, -1}
				continue
			}
			for col < len(grid[0])-1 {
				col++
				if grid[row][col] == 1 {
					continue
				}
				result[idx] = []int{row, col}
				break
			}
		}
		if string(query[0]) == "<" {
			row := toInt(string(query[2]))
			col := toInt(string(query[4]))
			if row-1 < len(grid[0])-1 || col-1 < len(grid)-1 {
				result[idx] = []int{-1, -1}
				continue
			}
			for row > 0 {
				row--
				if grid[col][row] == 1 {
					result[idx] = []int{-1, -1}
					continue
				}
				result[idx] = []int{row, col}
				break
			}
		}
	}

	for idx := 0; idx < len(result)-1; idx++ {
		if len(result[idx]) == 0 {
			result = slices.Delete(result, idx, idx+1)
		}
	}

	return result
}
