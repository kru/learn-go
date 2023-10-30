package core

import "fmt"

// https://www.hackerrank.com/challenges/one-month-preparation-kit-new-year-chaos/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-three
func MinimumBribes(queue []int32) {
	bribe := make(map[int32]int)

	for i := 0; i < len(queue); i++ {
		for j := i + 1; j <= len(queue)-1; j++ {
			if queue[i] > queue[j] {
				bribe[queue[i]] += 1
				queue[i], queue[j] = queue[j], queue[i]
			}
		}
	}

	fmt.Println(bribe)
	var op int
	for _, val := range bribe {
		if val > 2 {
			fmt.Println("Too chaotic")
			return
		} else {
			op += val
		}
	}

	fmt.Println(op)

}
