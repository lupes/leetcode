package question_451_460

import (
	"math"
)

// 456. 132模式
// https://leetcode-cn.com/problems/132-pattern
// Topics: 栈

func find132pattern(nums []int) bool {
	var stack, min, l = make([]int, 0), -math.MaxInt64, 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < min {
			return true
		}
		for l > 0 && nums[i] > stack[l-1] {
			stack, min, l = stack[:l-1], stack[l-1], l-1
		}
		stack, l = append(stack, nums[i]), l+1
	}
	return false
}
