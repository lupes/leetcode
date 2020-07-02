package question_441_450

import (
	"sort"
)

// 448. 找到所有数组中消失的数字
// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array
// Topics: 数组

func findDisappearedNumbers(nums []int) []int {
	nums = append(nums, len(nums)+1, 0)
	sort.Ints(nums)
	var res []int
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 > nums[i-1] {
			for j := nums[i-1] + 1; j < nums[i]; j++ {
				res = append(res, j)
			}
		}
	}
	return res
}
