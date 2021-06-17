package question_371_380

import (
	"container/heap"
)

// 373. 查找和最小的K对数字
// https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums
// Topics: 堆

type sumHeap [][]int

func (s sumHeap) Len() int {
	return len(s)
}

func (s sumHeap) Less(i, j int) bool {
	return (s[i][0]+s[i][1] < s[j][0]+s[j][1]) || (s[i][0]+s[i][1] == s[j][0]+s[j][1] && s[i][0] < s[j][0])
}

func (s sumHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *sumHeap) Push(x interface{}) {
	*s = append(*s, x.([]int))
}

func (s *sumHeap) Pop() interface{} {
	old := *s
	x := old[len(old)-1]
	*s = old[:len(old)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var h = &sumHeap{}
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			*h = append(*h, []int{n1, n2})
		}
	}

	if k >= len(nums1)*len(nums2) {
		return *h
	}

	heap.Init(h)

	var res [][]int
	for i := 0; i < k && i < len(nums1)*len(nums2); i++ {
		res = append(res, heap.Pop(h).([]int))
	}

	return res
}
