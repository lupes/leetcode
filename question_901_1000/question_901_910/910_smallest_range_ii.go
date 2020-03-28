package question_901_910

import (
	"sort"
)

// 910. 最小差值 II
// https://leetcode-cn.com/problems/smallest-range-ii
// Topics: 贪心算法 数学

func smallestRangeII(A []int, K int) int {
	sort.Ints(A)
	res, left, right := A[len(A)-1]-A[0], A[0], A[len(A)-1]
	for i := 0; i < len(A)-1; i++ {
		a, b := A[i], A[i+1]
		high, low := 0, 0
		if right-K > a+K {
			high = right - K
		} else {
			high = a + K
		}
		if left+K > b-K {
			low = b - K
		} else {
			low = left + K
		}
		if high-low < res {
			res = high - low
		}
	}
	return res
}
