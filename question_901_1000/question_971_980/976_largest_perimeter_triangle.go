package question_971_980

import (
	"sort"
)

// 976. 三角形的最大周长
// https://leetcode-cn.com/problems/largest-perimeter-triangle
// Topics: 排序 数学

func largestPerimeter(A []int) int {
	sort.Ints(A)
	var max = 0
	for i := len(A) - 1; i >= 2; i-- {
		if A[i-2]+A[i-1] > A[i] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return max
}
