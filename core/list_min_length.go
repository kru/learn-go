package core

func ListMinLength(list []int32, k int32) int32 {
	for i := 0; i < len(list)-1; i++ {
		tmp := list[i] * list[i+1]
		if tmp > k {
			continue
		}
		list[i+1] = tmp
		list = append(list[:i], list[i+1:]...)
		i = i - 1
	}
	return int32(len(list))
}
