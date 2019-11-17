package question_221_230

import (
	"fmt"
)

// 228. 汇总区间
// https://leetcode-cn.com/problems/summary-ranges/
// Topics: 数组

func summaryRanges(nums []int) []string {
	var res []string
	var i, j int
	for i = 0; i < len(nums); i = j + 1 {
		for j = i; j < len(nums)-1; j++ {
			if nums[j]+1 != nums[j+1] {
				break
			}
		}
		if i == j {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
	}
	return res
}
