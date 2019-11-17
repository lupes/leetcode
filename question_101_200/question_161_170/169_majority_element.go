package question_161_170

import "sort"

// 169. 求众数
// https://leetcode-cn.com/problems/majority-element
// Topics: 位运算 数组 分治算法

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
