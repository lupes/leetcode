package question_31_40

import (
	"sort"
)

// 39. 组合总和
// https://leetcode-cn.com/problems/combination-sum

func combinationSum(candidates []int, target int) [][]int {
	count := len(candidates)
	if count == 0 {
		return nil
	}
	sort.Ints(candidates)
	var res [][]int
	find(&res, nil, candidates, target, 0)
	return res
}

func find(res *[][]int, tmp, candidates []int, target, index int) {
	if target < candidates[0] {
		return
	}
	for i := index; i < len(candidates) && target >= candidates[i]; i++ {
		t := make([]int, len(tmp))
		copy(t, tmp)
		t = append(t, candidates[i])
		if target-candidates[i] == 0 {
			*res = append(*res, t)
			continue
		}
		find(res, t, candidates, target-candidates[i], i)
	}
}
