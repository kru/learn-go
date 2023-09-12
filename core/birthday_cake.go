package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kru/learn-go/core/util"
)

func countCandles(candles []int32) int32 {
	var total int32

	hInt := int32(0)
	for i := 0; i < len(candles); i++ {
		if candles[i] > hInt {
			hInt = candles[i]
		}
	}

	for j := 0; j < len(candles); j++ {
		if candles[j] == hInt {
			total += 1
		}
	}

	return total
}

func BirthdayCakeCandles() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	err := os.Setenv("OUTPUT_PATH", "core/test")
	if err != nil {
		return
	}
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	util.CheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	candlesCount, err := strconv.ParseInt(strings.TrimSpace(util.ReadLine(reader)), 10, 64)

	util.CheckError(err)

	candlesTemp := strings.Split(strings.TrimSpace(util.ReadLine(reader)), " ")

	var candles []int32

	for i := 0; i < int(candlesCount); i++ {
		candlesItemTemp, err := strconv.ParseInt(candlesTemp[i], 10, 64)
		util.CheckError(err)
		candlesItem := int32(candlesItemTemp)
		candles = append(candles, candlesItem)
	}

	result := countCandles(candles)

	fmt.Printf("candles: %d\n", result)

	err = writer.Flush()
	if err != nil {
		return
	}
}
