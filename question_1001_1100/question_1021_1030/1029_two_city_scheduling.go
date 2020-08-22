package question_1021_1030

import (
	"sort"
)

// 1029. 两地调度
// https://leetcode-cn.com/problems/two-city-scheduling
// Topics: 贪心算法

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		a := costs[i][0] - costs[i][1]
		if a < 0 {
			a *= -1
		}
		b := costs[j][0] - costs[j][1]
		if b < 0 {
			b *= -1
		}
		return a > b
	})
	a, b, l, r := 0, 0, len(costs)/2, 0
	for _, cost := range costs {
		if a == l || (b < l && cost[0] > cost[1]) {
			r += cost[1]
			b++
		} else {
			r += cost[0]
			a++
		}
	}
	return r
}
