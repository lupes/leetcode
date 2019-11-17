package question_171_180

import (
	"sort"
	"strconv"
)

// 179. 最大数
// https://leetcode-cn.com/problems/largest-number/
// Topics: 排序

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		if a+b < b+a {
			return false
		}
		return true
	})
	if len(nums) >= 1 && nums[0] == 0 {
		return "0"
	}
	var res = ""
	for _, n := range nums {
		res += strconv.Itoa(n)
	}
	return res
}
