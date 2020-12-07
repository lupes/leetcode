package question_31_40

import (
	"sort"
)

// 31. 下一个排列
// https://leetcode-cn.com/problems/next-permutation
// Topics: 数组

func nextPermutation(nums []int) {
	size := len(nums)

	for i := size - 2; i >= 0; i-- {
		min, index := 101, -1
		for j := i + 1; j < size; j++ {
			if nums[j] > nums[i] && nums[j] < min {
				min, index = nums[j], j
			}
		}
		if index != -1 {
			nums[i], nums[index] = nums[index], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}

	sort.Ints(nums)
}
