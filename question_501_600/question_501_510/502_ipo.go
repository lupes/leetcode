package question_501_510

import (
	"sort"
)

// 502. IPO
// https://leetcode-cn.com/problems/ipo/
// Topics: 堆 贪心算法

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	var projects, available = make([][2]int, 0, len(Profits)/2), make([][2]int, 0, len(Profits)/2)
	for i, c := range Capital {
		if c <= W {
			available = append(available, [2]int{Profits[i], Capital[i]})
		} else {
			projects = append(projects, [2]int{Profits[i], Capital[i]})
		}
	}
	sort.Slice(available, func(i, j int) bool {
		return available[i][0] > available[j][0]
	})

	sort.Slice(projects, func(i, j int) bool {
		return projects[i][1] < projects[j][1]
	})

	for ; k > 0 && len(available) > 0; k-- {
		if len(projects) == 0 {
			for ; k > 0 && len(available) > 0; k-- {
				W, available = W+available[0][0], available[1:]
			}
			return W
		}

		W, available = W+available[0][0], available[1:]
		for len(projects) > 0 && projects[0][1] <= W {
			available = append(available, projects[0])
			for k := len(available) - 1; k > 0 && available[k][0] > available[k-1][0]; k-- {
				available[k], available[k-1] = available[k-1], available[k]
			}
			projects = projects[1:]
		}
	}

	return W
}
