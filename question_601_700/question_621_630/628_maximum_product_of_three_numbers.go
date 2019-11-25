package question_621_630

import (
	"sort"
)

// 628. 三个数的最大乘积
// https://leetcode-cn.com/problems/maximum-product-of-three-numbers
// Topics: 数组 数学

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	res1 := nums[l-1] * nums[l-2] * nums[l-3]
	res2 := nums[l-1] * nums[0] * nums[1]
	if res1 < res2 {
		res1 = res2
	}
	return res1
}
