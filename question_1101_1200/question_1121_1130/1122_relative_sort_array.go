package question_1121_1130

import (
	"sort"
)

// 1122. 数组的相对排序
// https://leetcode-cn.com/problems/relative-sort-array
// Topics: 排序 数组

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var m = make(map[int][]int)
	for _, n := range arr1 {
		m[n] = append(m[n], n)
	}
	var res []int
	for _, n := range arr2 {
		res = append(res, m[n]...)
		delete(m, n)
	}
	var res2 []int
	for _, v := range m {
		res2 = append(res2, v...)
	}
	sort.Ints(res2)
	return append(res, res2...)
}
