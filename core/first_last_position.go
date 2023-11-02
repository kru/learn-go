package core

import "fmt"

func SearchRange(nums []int, target int) []int {
	end := len(nums) - 1
	for i := 0; i <= end; i++ {
		if target == nums[i] {
			nums = append(nums, i)
		}
	}
	fmt.Println(nums)
	if (len(nums)-1)-end == 0 {
		return []int{-1, -1}
	}
	if len(nums[end+1:]) == 1 {
		nums = append(nums[end+1:], nums[end+1])
	}
	if len(nums) > 2 {
		nums = nums[end+1:]
		nums = append(nums[:1], nums[len(nums)-1])
	}
	return nums
}
