package question_1601_1610

// 1608. 特殊数组的特征值
// https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
// Topics: 数组

import (
	"sort"
)

func specialArray(nums []int) int {
	nums = append(nums, -1)
	sort.Ints(nums)
	l := len(nums)
	for i := 1; i < len(nums); {
		t := l - i
		if t > nums[i-1] && nums[i] >= t {
			return t
		}
		j := i + 1
		for ; j < len(nums) && nums[i] == nums[j]; j++ {
		}
		i = j
	}
	return -1
}
