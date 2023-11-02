package core

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kru/learn-go/core/util"
)

// https://www.hackerrank.com/challenges/one-month-preparation-kit-new-year-chaos/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-three
// This didn't pass test case input06 and input08 respectively, will revisit later
func MinimumBribes(queue []int32) int32 {
	st := time.Now().Nanosecond()
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
				}
			}
		}
	}
	et := time.Now().Nanosecond()
	fmt.Printf("%d -> it tooks %.2f milisec\n", op, float64(et-st)/math.Pow10(6))
	return op
}

func InitBribe() {
	f, err := os.Open("core/custom-input/minimum-bribe-input08")
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

		MinimumBribes(q)

		f.Close()
	}
}
