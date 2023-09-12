package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kru/learn-go/core/util"
)

func Staircase() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(util.ReadLine(reader)), 10, 64)

	util.CheckError(err)
	n := int32(nTemp)
	stair := strings.Repeat("#", int(n))

	for j := 1; j < int(n+1); j++ {
		fmt.Println(strings.Replace(stair, "#", " ", int(n)-j))
	}
}
