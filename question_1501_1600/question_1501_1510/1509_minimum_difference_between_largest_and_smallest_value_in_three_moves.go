package question_1501_1510

// 1509. 三次操作后最大值与最小值的最小差
// https://leetcode-cn.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
// Topics: 排序 数组

import (
	"sort"
)

func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	sort.Ints(nums)
	left, right, min := 0, len(nums)-4, nums[len(nums)-1]-nums[0]
	for i := 0; i <= 3; i++ {
		if nums[right+i]-nums[left+i] < min {
			min = nums[right+i] - nums[left+i]
		}
	}
	return min
}
