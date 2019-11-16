package question_161_170

import "sort"

// 169. 求众数
// https://leetcode-cn.com/problems/majority-element

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
