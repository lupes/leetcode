package question_491_500

import (
	"reflect"
)

// 491. 递增子序列
// https://leetcode-cn.com/problems/increasing-subsequences/
// Topics: 深度优先搜索

func findSubsequences(nums []int) [][]int {
	if len(nums) < 2 {
		return nil
	}
	tmp := findSubSequencesDfs(nums)[1:]
	var res [][]int
	for _, t := range tmp {
		for _, r := range res {
			if reflect.DeepEqual(t, r) {
				goto Next
			}
		}
		res = append(res, t)
	Next:
	}
	return res
}

func findSubSequencesDfs(nums []int) [][]int {
	var res [][]int
	if len(nums) > 0 {
		res = append(res, []int{nums[0]})
	}
	if len(nums) < 2 {
		return res
	}

	var now = nums[0]
	for i, n := range nums[1:] {
		if n >= now {
			tmp := findSubSequencesDfs(nums[i+1:])
			for _, t := range tmp {
				if len(t) > 1 {
					res = append(res, t)
				}
				res = append(res, append([]int{now}, t...))
			}
		}
	}
	return res
}
