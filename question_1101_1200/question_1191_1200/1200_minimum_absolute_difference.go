package question_1191_1200

import (
	"sort"
)

// 1200. 最小绝对差
// https://leetcode-cn.com/problems/minimum-absolute-difference
// Topics: 数组

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var min = int(1e7)
	var res [][]int
	for i := 1; i < len(arr); i++ {
		a := arr[i] - arr[i-1]
		if a < 0 {
			a = -a
		}
		if a < min {
			min = a
			res = [][]int{{arr[i-1], arr[i]}}
		} else if a == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
