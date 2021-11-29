package question_781_790

import (
	"sort"
)

// 786. 第 K 个最小的素数分数
// https://leetcode-cn.com/problems/k-th-smallest-prime-fraction
// Topics: 堆 二分查找

func kthSmallestPrimeFraction(A []int, K int) []int {
	var tmp [][]int
	for i, a := range A {
		for j := i + 1; j < len(A); j++ {
			tmp = append(tmp, []int{a, A[j]})
		}
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][0]*tmp[j][1] < tmp[j][0]*tmp[i][1]
	})
	return tmp[K-1]
}
