package question_641_650

import (
	"sort"
)

// 646. 最长数对链
// https://leetcode-cn.com/problems/maximum-length-of-pair-chain
// Topics: 动态规划

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	var flag = make([]int, len(pairs))
	for i, pair := range pairs {
		var max = 0
		for j := i - 1; j >= 0; j-- {
			if pair[0] > pairs[j][1] && flag[j] > max {
				max = flag[j]
			}
		}
		flag[i] = max + 1
	}
	var max = 1
	for _, f := range flag {
		if f > max {
			max = f
		}
	}
	return max
}
