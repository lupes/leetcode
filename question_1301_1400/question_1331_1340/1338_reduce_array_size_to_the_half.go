package question_1311_1320

import (
	"sort"
)

// 1338. 数组大小减半
// https://leetcode-cn.com/problems/reduce-array-size-to-the-half/
// Topics: 贪心算法 数组

func minSetSize(arr []int) int {
	var m = make(map[int]int)
	for _, n := range arr {
		m[n]++
	}

	var flag = make([]int, 0, len(m))
	for _, v := range m {
		flag = append(flag, v)
	}

	sort.Slice(flag, func(i, j int) bool {
		return flag[i] > flag[j]
	})

	var l, t = len(arr), 0
	for i, n := range flag {
		t += n
		if t<<1 >= l {
			return i + 1
		}
	}
	return 0
}
