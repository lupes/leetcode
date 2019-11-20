package question_451_460

import (
	"sort"
)

// 455. 分发饼干
// https://leetcode-cn.com/problems/assign-cookies
// Topics: 贪心算法

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	var res int
	for gi, si := 0, 0; gi < len(g) && si < len(s); {
		if g[gi] > s[si] {
			gi++
		} else if g[gi] <= s[si] {
			si++
			gi++
			res++
		}
	}
	return res
}
