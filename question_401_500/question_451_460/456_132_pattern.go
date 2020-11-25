package question_451_460

import (
	"math"
)

// 456. 132模式
// https://leetcode-cn.com/problems/132-pattern
// Topics: 栈

func find132pattern(nums []int) bool {
	var stack []int
	min := -math.MaxInt64
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < min {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
