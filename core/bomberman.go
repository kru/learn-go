package core

import "fmt"

var direction = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func BomberMan(n int32, grid []string) []string {
	grid01 := make([][]rune, len(grid))

	if n == 1 {
		return grid
	}
	// see if we need to simulate the second using the j % 2 == 0
	for j := 2; j <= int(n); j++ {
		if j%2 == 0 {
			for i := 0; i < len(grid01); i++ {
				grid01[i] = make([]rune, len(grid[i]))

				for j := range grid01[i] {
					grid01[i][j] = 'O'
				}
			}
		} else {
			for row := range grid {
				for col := range grid[row] {
					if string(grid[row][col]) != "O" {
						continue
					}
					grid01[row][col] = '.'
					for _, dir := range direction {
						x, y := row+dir[0], col+dir[1]
						if x > len(grid)-1 || y > len(grid[row])-1 || x < 0 || y < 0 {
							continue
						}
						grid01[x][y] = '.'
					}
				}
			}
			for id, gr := range grid01 {
				s := ""
				for idx := range gr {
					s = s + string(gr[idx])
				}
				grid[id] = s
			}
		}
	}

	for id, gr := range grid01 {
		s := ""
		for idx := range gr {
			s = s + string(gr[idx])
		}
		grid[id] = s
	}
	fmt.Println(grid)
	return grid
}
