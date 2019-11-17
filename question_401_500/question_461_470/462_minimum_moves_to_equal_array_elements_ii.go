package question_461_470

import "sort"

// 462. 最少移动次数使数组元素相等 II
// https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/
// Topics: 数学

func minMoves2(nums []int) int {
	var l, r = len(nums), 0
	sort.Ints(nums)
	for i, j := 0, l-1; j > i; i, j = i+1, j-1 {
		r += nums[j] - nums[i]
	}
	return r
}
