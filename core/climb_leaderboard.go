package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kru/learn-go/core/util"
)

// Use binary search approach with opposite order from end to start of the slice
// @TODO: compare the performance versus sliding window approach
func findPosition(rank []int32, score int32) int32 {
	var prev int32

	if score >= rank[0] {
		return 1
	}

	if score < rank[len(rank)-1] {
		return int32(len(rank)) + 1
	}
	if score == rank[len(rank)-1] {
		return int32(len(rank))
	}

	var start, end int

	idx := len(rank) / 2

	if score > rank[idx] {
		start = idx + 1
		end = 0
	} else {
		start = len(rank)
		end = idx - 1
	}

	for i := start - 1; i >= end; i-- {

		if score == rank[i] {
			return int32(i) + 1
		}

		if rank[i] > score && score > prev {
			return int32(i + 2)
		}
		prev = rank[i]
	}

	// Impossible to happen
	return 0
}

// https://www.hackerrank.com/challenges/one-month-preparation-kit-climbing-the-leaderboard/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-three
func ClimbLeaderBoard(ranked, player []int32) []int32 {
	st := time.Now()
	var pos, rank []int32
	occr := make(map[int32]int)
	for i := range ranked {
		if occr[ranked[i]] == 0 {
			rank = append(rank, ranked[i])
		}
		occr[ranked[i]] += 1
	}

	for _, score := range player {
		res := findPosition(rank, score)
		pos = append(pos, res)
	}
	fmt.Println(pos)
	fmt.Println(time.Since(st).Milliseconds(), "ms")
	return pos
}

// Big input
func InitClimbLeaderboard() {
	f, err := os.Open("core/custom-input/input06-climb-leaderboard")
	util.CheckError(err)

	reader := bufio.NewReaderSize(f, 16*1024*1024)

	err = os.Setenv("OUTPUT_PATH", "core/test")
	if err != nil {
		return
	}
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	util.CheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(util.ReadLine(reader)), 10, 64)
	util.CheckError(err)

	rankedTemp := strings.Split(strings.TrimSpace(util.ReadLine(reader)), " ")

	var ranked []int32
	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		util.CheckError(err)

		ranked = append(ranked, int32(rankedItemTemp))
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(util.ReadLine(reader)), 10, 64)
	util.CheckError(err)

	playerTemp := strings.Split(strings.TrimSpace(util.ReadLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		util.CheckError(err)

		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := ClimbLeaderBoard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}
