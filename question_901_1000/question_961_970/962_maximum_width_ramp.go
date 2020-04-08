package question_961_970

import (
	"sort"
)

// 962. 最大宽度坡
// https://leetcode-cn.com/problems/maximum-width-ramp
// Topics: 数组

func maxWidthRamp(A []int) int {
	var B = make([][2]int, len(A))
	for i := range A {
		B[i] = [2]int{i, A[i]}
	}
	sort.Slice(B, func(i, j int) bool {
		if B[i][1] < B[j][1] {
			return true
		} else if B[i][1] == B[j][1] && B[i][0] < B[j][0] {
			return true
		}
		return false
	})
	m, res := len(A), 0
	for _, t := range B {
		if t[0]-m > res {
			res = t[0] - m
		}
		if m > t[0] {
			m = t[0]
		}
	}
	return res
}
