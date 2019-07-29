package question_161_170

import "sort"

// 164. 最大间距
//  https://leetcode-cn.com/problems/maximum-gap/

func maximumGap(nums []int) int {
	sort.Ints(nums)
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > max {
			max = nums[i+1] - nums[i]
		}
	}
	return max
}
