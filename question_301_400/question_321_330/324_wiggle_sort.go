package question_321_330

import (
	"sort"
)

// 324. 摆动排序 II
// https://leetcode-cn.com/problems/wiggle-sort-ii/

func wiggleSort(nums []int) {
	sort.Ints(nums)
	l := len(nums)
	if l < 3 {
		return
	}
	c := l / 2
	r := l - 1
	if l%2 != 0 {
		r--
	}
	var res []int
	for i := c - 1; i >= 0; {
		res = append(res, nums[i], nums[r])
		r--
		i--
	}
	if l%2 != 0 {
		res = append(res, nums[l-1])
		res[l-1], res[l-2] = res[l-2], res[l-1]
	}

	copy(nums, res)
}
