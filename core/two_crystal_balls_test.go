package core

import "testing"

type crystalBall struct {
	input    []int
	expected int
}

var crystalBalls = []crystalBall{
	{[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1}, 8},
	{[]int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, 4},
	{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1}, 9},
	{[]int{0, 0, 0, 0, 1, 1}, 5},
}

func TestTwoCrystalBall(t *testing.T) {

	for _, test := range crystalBalls {
		if got := TwoCrystalBall(test.input); got != test.expected {
			t.Errorf("got %d, test.expected %d", got, test.expected)
		}
	}
}

func BenchmarkTwoCrystalBall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoCrystalBall([]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1})
	}
}
