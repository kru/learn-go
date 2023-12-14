package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kru/learn-go/core/util"
)

// answer 55614
func TrebuchetSpell() {
	f, err := os.Open("core/custom-input/input")
	util.CheckError(err)

	fScanner := bufio.NewScanner(f)
	fScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fScanner.Scan() {
		fileLines = append(fileLines, fScanner.Text())
	}

	f.Close()

	var total int64
	for _, line := range fileLines {
		firstLast := ""

		str3 := ""
		str4 := ""
		str5 := ""
		for i := 0; i < len(line); i++ {

			num, _ := strconv.ParseInt(string(line[i]), 10, 64)
			if num > 0 {
				firstLast += string(line[i])
				continue
			}
			ln3 := 3
			ln4 := 4
			ln5 := 5

			if i+ln3 <= len(line) {
				str3 = line[i : i+ln3]
			}
			if i+ln4 <= len(line) {
				str4 = line[i : i+ln4]
			}
			if i+ln5 <= len(line) {
				str5 = line[i : i+ln5]
			}

			if len(str3) < 3 && len(str4) < 4 && len(str5) < 5 {
				continue
			}

			if len(str3) == 3 {

				if str3 == "one" {
					firstLast += "1"
					continue
				}
				if str3 == "two" {
					firstLast += "2"
					continue
				}
				if str3 == "six" {
					firstLast += "6"
					continue
				}
				str3 = str3[1:]
			}

			if len(str4) == 4 {
				if str4 == "four" {
					firstLast += "4"
					continue
				}
				if str4 == "five" {
					firstLast += "5"
					continue
				}
				if str4 == "nine" {
					firstLast += "9"
					continue
				}
				str4 = str4[1:]

			}

			if len(str5) == 5 {
				if str5 == "three" {
					firstLast += "3"
					continue
				}
				if str5 == "seven" {
					firstLast += "7"
					continue
				}
				if str5 == "eight" {
					firstLast += "8"
					continue
				}
				str5 = str5[1:]
			}
		}

		fl, _ := strconv.ParseInt(string(firstLast[0])+string(firstLast[len(firstLast)-1]), 10, 64)

		total += fl
	}

	fmt.Println()
	fmt.Println(total)
}
