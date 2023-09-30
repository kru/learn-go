package core

import (
	"fmt"
	"strconv"

	"github.com/kru/learn-go/core/util"
)

func recursiveSum(n string, total int32) int32 {
	if len(n) == 0 {
		s := strconv.Itoa(int(total))
		if len(s) == 1 {
			fmt.Println(total)
			return total
		} else {
			return recursiveSum(s, 0)
		}
	}

	num, err := strconv.Atoi(string(n[0]))
	util.CheckError(err)

	total += int32(num)

	return recursiveSum(n[1:], total)
}

func SuperDigit(n string, k int32) int32 {
	if len(n) == 1 {
		r, err := strconv.Atoi(n)
		util.CheckError(err)
		return int32(r)
	}

	var m int32
	for j := 0; j < len(n); j++ {
		in, err := strconv.Atoi(string(n[j]))
		util.CheckError(err)
		m += int32(in)
	}

	m = m * k
	fmt.Println(m)
	return recursiveSum(strconv.Itoa(int(m)), 0)
}
