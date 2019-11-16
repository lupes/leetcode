package question_31_40

import (
	"sort"
)

// 31. 下一个排列
// https://leetcode-cn.com/problems/next-permutation

func nextPermutation(nums []int) {
	size := len(nums)
	if size < 2 {
		return
	}

	var (
		min      int
		minIndex int
		big      int
		bigIndex int
	)

	for i := size - 1; i >= 0; i-- {
		if i == 0 {
			minIndex = -1
			break
		}
		if nums[i] > nums[i-1] {
			min = nums[i-1]
			minIndex = i - 1
			big = nums[i]
			bigIndex = i
			break
		}
	}
	if minIndex == -1 {
		sort.Ints(nums)
		return
	}

	for i := bigIndex + 1; i < size; i++ {
		if big > nums[i] && nums[i] > min {
			bigIndex = i
			big = nums[i]
		}
	}

	nums[minIndex], nums[bigIndex] = nums[bigIndex], nums[minIndex]
	sort.Ints(nums[minIndex+1:])
}
