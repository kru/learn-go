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
		[]string{
			".......",
			"...O.O.",
			"...OO..",
			"..OOOO.",
			"OOOOOOO",
			"OOOOOOO"},
	},
	{
		181054341,
		[]string{"O..OO........O..O........OO.O.OO.OO...O.....OOO...OO.O..OOOOO...O.O..O..O.O..OOO..O..O..O....O...O....O...O..O..O....O.O.O.O.....O.....OOOO..O......O.O.....OOO....OO....OO....O.O...O..OO....OO..O...O"},
		[]string{"OOOOO........OOOO........OOOOOOOOOO...O.....OOO...OOOOOOOOOOO...OOOOOOOOOOOOOOOOOOOOOOOOO....O...O....O...OOOOOOO....OOOOOOO.....O.....OOOOOOO......OOO.....OOO....OO....OO....OOO...OOOOO....OOOOO...O"},
	},
}

func TestBomberMan(t *testing.T) {
	for _, test := range payloads {
		result := BomberMan(test.time, test.grid)
		for i := range result {
			if result[i] != test.expected[i] {
				t.Fatalf("got %s, test expected %s", result[i], test.expected[i])
			}
		}
	}
}
