package question_1311_1320

import (
	"sort"
)

// 1331. 数组序号转换
// https://leetcode-cn.com/problems/rank-transform-of-an-array/
// Topics: 数组

func arrayRankTransform(arr []int) []int {
	var flag []int
	var m = make(map[int]int)
	for _, n := range arr {
		if _, ok := m[n]; !ok {
			flag = append(flag, n)
			m[n] = -1
		}
	}

	sort.Ints(flag)
	for i, n := range flag {
		m[n] = i + 1
	}

	for i, n := range arr {
		arr[i] = m[n]
	}
	return arr
}
