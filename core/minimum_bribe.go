package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kru/learn-go/core/util"
)

// https://www.hackerrank.com/challenges/one-month-preparation-kit-new-year-chaos/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-three
// This didn't pass test case input06 and input08, will revisit later
func MinimumBribe(queue []int32) int32 {
	var op int32
	for i := 0; i < len(queue)-1; i++ {
		diff := queue[i] - int32(i+1)
		if diff > 2 {
			fmt.Println("Too chaotic")
			return -1
		}

		if diff > 0 {
			op += diff
		} else if diff <= 0 {
			for j := i + 1; j <= len(queue)-1; j++ {
				if queue[i] > queue[j] {
					op += 1
					break
				}
			}
		}
	}
	// fmt.Printf("op %d, %v\n", op, queue)

	fmt.Println(op)
	return int32(op)
}

// expected output from test case 6
// 96110
// Too chaotic
// Too chaotic
// 96319
// 96323
// 96058

// expected output from test case 8
// 119814
// 120175
// 119810
// 119827
// Too chaotic

func InitBribe() {
	f, err := os.Open("core/custom-input/minimum-bribe-input06")
	util.CheckError(err)

	reader := bufio.NewReaderSize(f, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(util.ReadLine(reader)), 10, 64)
	util.CheckError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(util.ReadLine(reader)), 10, 64)
		util.CheckError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(util.ReadLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			util.CheckError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		MinimumBribe(q)

		f.Close()
	}
}
