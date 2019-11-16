package question_41_50

import (
	"sort"
)

// 41. 缺失的第一个正数
// https://leetcode-cn.com/problems/first-missing-positive

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	index := 1
	for i, n := range nums {
		if n > 0 {
			if i > 0 && n == nums[i-1] {
				continue
			}
			if n != index {
				return index
			}
			index++
		}
	}
	return index
}
