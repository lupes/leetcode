package question_951_960

import (
	"sort"
)

// 954. 二倍数对数组
// https://leetcode-cn.com/problems/array-of-doubled-pairs
// Topics: 数组 哈希表

func canReorderDoubled(A []int) bool {
	var flag = make(map[int]int)
	var tmp = make([]int, 0, len(A)/4)
	for _, n := range A {
		if _, ok := flag[n]; !ok {
			tmp = append(tmp, n)
		}
		flag[n]++
	}
	sort.Ints(tmp)

	for _, n := range tmp {
		if flag[n] == 0 {
			continue
		}
		t := n * 2
		if n < 0 {
			if n%2 == -1 {
				return false
			}
			t = n / 2
		}
		if flag[n] > flag[t] {
			return false
		}
		flag[t] -= flag[n]
	}
	return true
}
