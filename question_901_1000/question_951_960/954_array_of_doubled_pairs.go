package question_951_960

import (
	"sort"
)

// 954. 二倍数对数组
// https://leetcode-cn.com/problems/array-of-doubled-pairs
// Topics: 数组 哈希表

func canReorderDoubled(A []int) bool {
	var flag = make(map[int]int)
	var tmp []int
	for _, n := range A {
		if _, ok := flag[n]; !ok {
			tmp = append(tmp, n)
		}
		flag[n]++
	}
	sort.Ints(tmp)

	for i := 0; i < len(tmp); {
		n := tmp[i]
		if flag[n] == 0 {
			i++
			continue
		}
		n2 := n * 2
		if n < 0 {
			n2 = n / 2
		}
		t, ok := flag[n2]
		if ok && t > 0 {
			flag[n2]--
			flag[n]--
		} else {
			return false
		}
	}
	return true
}
