package question_131_140

import (
	"sort"
)

// 136. 只出现一次的数字
// https://leetcode-cn.com/problems/single-number

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		if i == len(nums)-1 {
			return nums[i]
		} else if nums[i] != nums[i+1] {
			return nums[i]
		}
		i += 2
	}
	return nums[0]
}
