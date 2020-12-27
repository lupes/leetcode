package question_941_950

import (
	"sort"
)

// 945. 使数组唯一的最小增量
// https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique
// Topics: 数组

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	res, tmp := 0, 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			tmp = A[i-1] - A[i] + 1
			A[i] += tmp
			res += tmp
		}
	}
	return res
}
