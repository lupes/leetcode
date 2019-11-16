package question_31_40

import (
	"sort"
)

// 40. 组合总和 II
// https://leetcode-cn.com/problems/combination-sum-ii

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs2(candidates, target)
}

func dfs2(candidates []int, target int) [][]int {
	if candidates == nil {
		return nil
	}
	var res [][]int
	for i, n := range candidates {
		if i > 0 && n == candidates[i-1] {
			continue
		}
		if target == n {
			res = append(res, []int{n})
		} else if target > n {
			tmp := dfs2(candidates[i+1:], target-n)
			for _, t := range tmp {
				r := append([]int{n}, t...)
				res = append(res, r)
			}
		}
	}
	return res
}
