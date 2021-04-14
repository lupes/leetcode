package question_1611_1620

// 1619. 删除某些元素后的数组均值
// https://leetcode-cn.com/problems/mean-of-array-after-removing-some-elements
// Topics: 数组

import (
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	t, l, sum := len(arr)/20, len(arr), 0
	for i := t; i < l-t; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(l-2*t)
}
