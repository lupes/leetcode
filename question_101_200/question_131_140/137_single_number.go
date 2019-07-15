package question_131_140

import (
	"sort"
)

// 137. 只出现一次的数字 II
// https://leetcode-cn.com/problems/single-number-ii/

func singleNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i += 3 {
		if nums[i-1] != nums[i] {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}
