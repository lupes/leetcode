package question_271_280

import "sort"

// 274. H指数
// https://leetcode-cn.com/problems/h-index/
// Topics: 排序 哈希表

func hIndex(citations []int) int {
	sort.Ints(citations)
	l := len(citations)
	for i, n := range citations {
		if n >= l-i {
			return l - i
		}
	}
	return 0
}
