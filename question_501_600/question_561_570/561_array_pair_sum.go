package question_561_570

import "sort"

// 561. 数组拆分 I
// https://leetcode-cn.com/problems/array-partition-i/

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}
