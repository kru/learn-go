package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func removeFormat(s string) string {
	s = strings.Replace(s, "PM", "", 1)
	return strings.Replace(s, "AM", "", 1)
}

func exist(s string, ante string) bool {
	return strings.Contains(strings.ToUpper(s), ante)
}

//Note:
//- 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
//- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
//- 11:59PM => 23:59

func Convert(s string) string {
	list := strings.SplitAfter(s, ":")
	hh := strings.Replace(string(list[0]), ":", "", 1)

	var h string
	num, _ := strconv.ParseInt(fmt.Sprint(strings.TrimSpace(hh)), 10, 0)

	if num == 0 {
		return removeFormat(s)
	}

	if num > 00 && num < 12 && exist(s, "PM") {
		num = num + 12
		h = fmt.Sprint(num)

		s = strings.Replace(s, hh, h, 1)
		return removeFormat(s)
	}

	if num == 12 && exist(s, "AM") {
		h = fmt.Sprint(0, "0")

		s = strings.Replace(s, hh, h, 1)
		return removeFormat(s)
	}

	if num < 10 {
		h = fmt.Sprint("0", num)
	} else {
		h = fmt.Sprint(num)
	}

	s = strings.Replace(s, hh, h, 1)
	return removeFormat(s)
}

func TimeConversion() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	err := os.Setenv("OUTPUT_PATH", "core/test")
	if err != nil {
		return
	}
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	CheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := ReadLine(reader)

	result := Convert(s)
	fmt.Printf("result %s\n", result)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}
