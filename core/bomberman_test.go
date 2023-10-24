package core

import "testing"

type payload struct {
	time     int32
	grid     []string
	expected []string
}

var payloads = []payload{
	{
		3,
		[]string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."},
		[]string{"OOO.OOO", "OO...OO", "OOO...O", "..OO.OO", "...OOOO", "...OOOO"},
	},
	{
		5,
		[]string{".......", "...O.O.", "....O..", "..O....", "OO...OO", "OO.O..."},
		[]string{".......", "...O.O.", "...OO..", "..OOOO.", "OOOOOOO", "OOOOOOO"},
	},
	{
		3,
		[]string{"....O..O."},
		[]string{"OOO......"},
	},
	{
		5,
		[]string{"....O..O."},
		[]string{"....OOOOO"},
	},
	{
		181054341,
		[]string{"O..OO........O..O........OO.O.OO.OO...O.....OOO...OO.O..OOOOO...O.O..O..O.O..OOO..O..O..O....O...O....O...O..O..O....O.O.O.O.....O.....OOOO..O......O.O.....OOO....OO....OO....O.O...O..OO....OO..O...O"},
		[]string{"OOOOO........OOOO........OOOOOOOOOO...O.....OOO...OOOOOOOOOOO...OOOOOOOOOOOOOOOOOOOOOOOOO....O...O....O...OOOOOOO....OOOOOOO.....O.....OOOOOOO......OOO.....OOO....OO....OO....OOO...OOOOO....OOOOO...O"},
	},
	{
		329973043,
		[]string{"OOOO.O.O...OOO.O.O........O.OOO.O.....OO..O..O...OOO....O.OOO....O...O....O..O.O.O.....OOOO.O...O....OO.O...........O.O..O.O..O...OO.OOO......O........O...O....O.O..O....O.......OOOO.O..........OO.O"},
		[]string{".........O.........OOOOOO.........OOO..........O.....OO.......OO...O...OO..........OOO........O...OO......OOOOOOOOO.............O........OOOO...OOOOOO...O...OO........OO...OOOOO........OOOOOOOO....."},
	},
	{
		4,
		[]string{"..O...O"},
		[]string{"OOOOOOO"},
	},
	{
		3,
		[]string{"O", ".", ".", ".", "O"},
		[]string{".", ".", "O", ".", "."},
	},
}

func TestBomberMan(t *testing.T) {
	for _, test := range payloads {
		result := BomberMan(test.time, test.grid)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Fatalf("got %s\n, test expected %s\n", result[i], test.expected[i])
			}
		}
	}
}
