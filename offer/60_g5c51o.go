package offer

import (
	"sort"
)

// 剑指 Offer II 060. 出现频率最高的 k 个数字
// https://leetcode-cn.com/problems/g5c51o/
// Topics: 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列）

func topKFrequent(nums []int, k int) []int {
	var flag1 = make(map[int]int)
	for _, n := range nums {
		flag1[n]++
	}

	var flag2 = make(map[int][]int)
	var keys []int
	for n, c := range flag1 {
		_, ok := flag2[c]
		if !ok {
			keys = append(keys, c)
		}
		flag2[c] = append(flag2[c], n)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	var res []int
	for i := 0; len(res) < k; i++ {
		res = append(res, flag2[keys[i]]...)
	}
	return res
}
