package question_1461_1470

// 1464. 数组中两元素的最大乘积
// https://leetcode-cn.com/problems/maximum-product-of-two-elements-in-an-array/
// Topics: 数组

import (
	"sort"
)

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
