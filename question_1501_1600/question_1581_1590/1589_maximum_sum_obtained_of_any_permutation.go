package question_1581_1590

import (
	"sort"
)

// 1589. 所有排列中的最大和
// https://leetcode-cn.com/problems/maximum-sum-obtained-of-any-permutation/
// Topics: 贪心算法

func maxSumRangeQuery(nums []int, requests [][]int) int {
	var diff, flag = make([]int, len(nums)+1), make([]int, len(nums))
	for _, req := range requests {
		diff[req[0]] += 1
		diff[req[1]+1] -= 1
	}

	flag[0] = diff[0]
	for i := 1; i < len(flag); i++ {
		flag[i] = flag[i-1] + diff[i]
	}

	sort.Ints(nums)
	sort.Ints(flag)
	var res = 0
	for i, n := range nums {
		res = (res + n*flag[i]) % (1e9 + 7)
	}
	return res
}
