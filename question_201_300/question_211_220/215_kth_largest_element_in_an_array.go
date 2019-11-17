package question_211_220

import "sort"

// 215. 数组中的第K个最大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
