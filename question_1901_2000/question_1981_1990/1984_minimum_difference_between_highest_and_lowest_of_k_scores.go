package question_1981_1990

import (
	"sort"
)

// 1984. 学生分数的最小差值
// https://leetcode-cn.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// Topic: 数组 排序 滑动窗口

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	l := len(nums)
	min := 1000000
	for t := k - 1; t < l; t++ {
		now := nums[t] - nums[t-k+1]
		if now < min {
			min = now
		}
	}
	return min
}
