package mian_shi

import (
	"sort"
)

// 面试题 17.14. 最小K个数
// https://leetcode-cn.com/problems/smallest-k-lcci/
// Topics: 数组 分治 快速选择 排序 堆

func smallestK(arr []int, k int) []int {
	if len(arr)/2 > k {
		a1 := smallestK(arr[:len(arr)/2], k)
		a2 := smallestK(arr[len(arr)/2:], k)
		arr = append(a1, a2...)
	}
	sort.Ints(arr)
	return arr[:k]
}
