package question_531_540

import (
	"sort"
)

// 532. 数组中的K-diff数对
// https://leetcode-cn.com/problems/k-diff-pairs-in-an-array
// Topics: 数组 双指针

func findPairs(nums []int, k int) int {
	var flag = make(map[int]int)
	var res int
	for _, n := range nums {
		flag[n]++
		if flag[n] == 2 && k == 0 {
			res++
		}
	}
	if k == 0 {
		return res
	}

	nums = make([]int, 0, len(flag))
	for k, _ := range flag {
		nums = append(nums, k)
	}
	sort.Ints(nums)

	for l, r := 0, 0; r >= l && r < len(nums); {
		if nums[r]-nums[l] == k {
			res++
			r++
		} else if nums[r]-nums[l] > k {
			l++
		} else {
			r++
		}
	}
	return res
}
