package core

import "fmt"

func DynamicList(n int32, queries [][]int32) []int32 {
	var lastAnswer int32
	lists := make([][]int32, n)

	var res []int32
	// query is [type, x, y]
	for _, query := range queries {
		if query[0] == 1 {
			idx := (query[1] ^ lastAnswer) % n

			lists[idx] = append(lists[idx], query[2])
		}
		if query[0] == 2 {
			id := (query[1] ^ lastAnswer) % n
			y := int(query[2]) % len(lists[id])
			res = append(res, lists[id][y])
			lastAnswer = lists[id][y]
		}
	}
	fmt.Println(res)
	return res
}
