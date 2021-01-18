package question_1401_1410

import (
	"sort"
)

// 1402. 做菜顺序
// https://leetcode-cn.com/problems/reducing-dishes/
// Topics: 动态规划

func maxSatisfaction(satisfaction []int) int {
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})

	var sum, res = 0, 0
	for _, n := range satisfaction {
		sum += n
		if res+sum > res {
			res = sum + res
		}
		if sum < 0 {
			break
		}
	}

	return res
}
