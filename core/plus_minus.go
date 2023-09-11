package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NumberRatio(arr []int32) {
	var po, ne, zr float32

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zr++
		} else if arr[i] > 0 {
			po++
		} else {
			ne++
		}
	}
	ln := float32(len(arr))
	fmt.Printf("%.6f\n", po/ln)
	fmt.Printf("%.6f\n", ne/ln)
	fmt.Printf("%.6f\n", zr/ln)
}

func PlusMinus() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(ReadLine(reader)), 10, 64)
	CheckError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(ReadLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		CheckError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	NumberRatio(arr)
}
