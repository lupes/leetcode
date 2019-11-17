package question_501_510

import (
	"sort"
)

// 502. IPO
// https://leetcode-cn.com/problems/ipo/
// Topics: 堆 贪心算法

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	var res = W
	var tmp = make(map[int][]int, 0)
	for i, c := range Capital {
		tmp[c] = append(tmp[c], Profits[i])
	}
	for k, v := range tmp {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
		tmp[k] = v
	}
	var stack = make(map[int][]int, 0)
	for k > 0 {
		for c, v := range tmp {
			if c <= res {
				stack[c] = v
				delete(tmp, c)
			}
		}
		max, index := 0, -1
		for c, v := range stack {
			if v[0] > max {
				max = v[0]
				index = c
			}
		}
		if index != -1 {
			res += max
			if len(stack[index]) == 1 {
				delete(stack, index)
			} else {
				stack[index] = stack[index][1:]
			}
			k--
		} else {
			return res
		}
	}
	return res
}
