package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumArr(idx int64, arr []int32) int64 {
	var res int64
	for i := 0; i < len(arr); i++ {
		if int64(i) != idx {
			res = res + int64(arr[i])
		}
	}
	return res
}

func Minmaxsum() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrTemp := strings.Split(strings.TrimSpace(ReadLine(reader)), " ")

	if len(arrTemp) != 5 {
		fmt.Println("Error: Input must be 5 integer space separated, e.g 1 2 3 6 9")
		return
	}
	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		CheckError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	var min, max int64

	for j := 0; j < len(arr); j++ {
		result := sumArr(int64(j), arr)
		if j == 0 {
			min = result
		}
		if result <= min {
			min = result
		}
		if result > max {
			max = result
		}
	}

	fmt.Printf("%v %v\n", min, max)
}
