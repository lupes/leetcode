package question_211_220

import (
	"sort"
)

// 220. 存在重复元素 III
// https://leetcode-cn.com/problems/contains-duplicate-iii/
// Topics: 排序 None

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	type N struct {
		n int
		i int
	}
	var tmp []N
	for i, n := range nums {
		tmp = append(tmp, N{
			n: n,
			i: i,
		})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].n < tmp[j].n
	})

	for i, n := range tmp {
		for j := i + 1; j < len(tmp) && n.n+t >= tmp[j].n; j++ {
			a := n.i - tmp[j].i
			if (a < 0 && -a <= k) || (a > 0 && k >= a) {
				return true
			}
		}
	}
	return false
}
