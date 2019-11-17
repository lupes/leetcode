package question_491_500

import (
	"sort"
)

// 502. IPO
// https://leetcode-cn.com/problems/ipo/
// Topics: 堆 贪心算法

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	var res = W
	var tmp = make(map[int][][2]int, 0)
	for i, c := range Capital {
		tmp[c] = append(tmp[c], [2]int{i, Profits[i]})
	}
	for k, v := range tmp {
		sort.Slice(v, func(i, j int) bool {
			return v[i][1] > v[j][1]
		})
		tmp[k] = v
	}
	for k > 0 {
		max, index := 0, -1
		for i := 0; i <= res; i++ {
			if arr, ok := tmp[i]; ok && max < arr[0][1] {
				max = arr[0][1]
				index = i
			}
		}
		if index != -1 {
			res += max
			if len(tmp[index]) == 1 {
				delete(tmp, index)
			} else {
				tmp[index] = tmp[index][1:]
			}
			k--
		} else {
			return res
		}
	}
	return res
}
