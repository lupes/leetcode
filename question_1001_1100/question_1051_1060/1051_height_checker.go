package question_1051_1060

import (
	"sort"
)

// 1051. 高度检查器
// https://leetcode-cn.com/problems/height-checker
// Topics: 数组

func heightChecker(heights []int) int {
	var tmp = make([]int, len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)
	var res = 0
	for i, n := range heights {
		if n != tmp[i] {
			res++
		}
	}
	return res
}
