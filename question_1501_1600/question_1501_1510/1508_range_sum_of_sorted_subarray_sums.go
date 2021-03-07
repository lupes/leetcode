package question_1501_1510

// 1508. 子数组和排序后的区间和
// https://leetcode-cn.com/problems/range-sum-of-sorted-subarray-sums/
// Topics: 排序 数组

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	var flag, tmp, res = make([]int, len(nums)+1), make([]int, 0, n*(n+1)/2), 0

	for i, t := range nums {
		flag[i+1] = flag[i] + t
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			tmp = append(tmp, flag[j]-flag[i])
		}
	}

	sort.Ints(tmp)
	for i := left - 1; i < right; i++ {
		res = (res + tmp[i]) % 1000000007
	}

	return res
}
