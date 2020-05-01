package question_1001_1010

import (
	"sort"
)

// 1005. K 次取反后最大化的数组和
// https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations
// Topics: 贪心算法

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	a, b, f, r, t := 101, 0, 0, 0, 0
	for i, n := range A {
		if n <= 0 {
			f = i + 1
		}
		if n >= 0 && n < a {
			a, b = n, n
		} else if n < 0 && n < a {
			a, b = -n, n
		}
	}
	if K > f {
		t, K = K-f, f
		if t%2 == 1 {
			if b > 0 {
				r -= 2 * b
			} else {
				r += 2 * b
			}
		}
	}
	for _, n := range A {
		if K > 0 {
			r += -n
			K--
		} else {
			r += n
		}
	}
	return r
}
