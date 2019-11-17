package question_341_350

import (
	"sort"
)

// 347. 前 K 个高频元素
// https://leetcode-cn.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) []int {
	var m = make(map[int]int)
	for _, i := range nums {
		m[i] += 1
	}
	var a [][]int
	for k, v := range m {
		a = append(a, []int{k, v})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] > a[j][1]
	})

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, a[i][0])
	}
	return res
}
