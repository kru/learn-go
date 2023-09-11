package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Staircase() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(ReadLine(reader)), 10, 64)

	CheckError(err)
	n := int32(nTemp)
	stair := strings.Repeat("#", int(n))

	for j := 1; j < int(n+1); j++ {
		fmt.Println(strings.Replace(stair, "#", " ", int(n)-j))
	}
}
