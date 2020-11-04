package question_1331_1340

import (
	"sort"
)

// 1356. 根据数字二进制下 1 的数目排序
// https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/
// Topics: 排序 位运算

type bit struct {
	n int
	b int
}

func sortByBits(arr []int) []int {
	var bits = make([]bit, len(arr))
	for i, n := range arr {
		var b = bit{n: n, b: 0}
		for ; n > 0; n /= 2 {
			if n%2 == 1 {
				b.b++
			}
		}
		bits[i] = b
	}

	sort.Slice(bits, func(i, j int) bool {
		return bits[i].b < bits[j].b || (bits[i].b == bits[j].b && bits[i].n < bits[j].n)
	})

	for i, b := range bits {
		arr[i] = b.n
	}
	return arr
}
