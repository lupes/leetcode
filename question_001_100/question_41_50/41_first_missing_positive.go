package question_41_50

import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	index := 1
	for i, n := range nums {
		if n > 0 {
			if i > 0 && n == nums[i-1] {
				continue
			}
			if n != index {
				return index
			}
			index++
		}
	}
	return index
}
