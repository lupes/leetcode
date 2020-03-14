package question_911_920

import (
	"sort"
)

// 912. 排序数组
// https://leetcode-cn.com/problems/sort-an-array
// Topics:

func sortArray(nums []int) []int {
	sort.Ints(nums)
	//var t, min = 0, 0
	//for i := 0; i < len(nums)-1; i++ {
	//	t, min = i, nums[i]
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[j] < min {
	//			min = nums[j]
	//			t = j
	//		}
	//	}
	//	if t != i {
	//		nums[i], nums[t] = nums[t], nums[i]
	//	}
	//}
	return nums
}
