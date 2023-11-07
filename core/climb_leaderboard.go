package core

import "fmt"

func findPosition(rank []int32, occr map[int32]int, score int32) int32 {

	var idx, ctr, prev int32

	if score > rank[0] {
		return 1
	}

	for i := 0; i < len(rank); i++ {
		if score == rank[i] {
			return int32(i) + 1
		}

		if score < prev && score > rank[i] {
			return int32(i) + 1
		}
		prev = rank[i]
	}
	if idx == 0 {
		idx = int32(len(rank)) + 1
	}

	fmt.Println(idx, ctr, occr)
	return idx
}

// https://www.hackerrank.com/challenges/one-month-preparation-kit-climbing-the-leaderboard/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-three
// let's use binary search to determine the position
func ClimbLeaderBoard(ranked, player []int32) []int32 {
	var pos, rank []int32
	occr := make(map[int32]int)
	for i := range ranked {
		if occr[ranked[i]] == 0 {
			rank = append(rank, ranked[i])
		}
		occr[ranked[i]] += 1
	}
	for _, score := range player {
		res := findPosition(rank, occr, score)
		pos = append(pos, res)
	}
	fmt.Println(pos)
	return pos
}
