package question_1491_1500

// 1498. 满足条件的子序列数目
// https://leetcode-cn.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
// Topics: 排序 滑动窗口

import (
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	left, right, res, mod, flag := 0, len(nums)-1, 0, 1000000007, make([]int, len(nums))
	flag[0] = 1
	for i := 1; i < len(nums); i++ {
		flag[i] = flag[i-1] * 2 % mod
	}
	for right >= left {
		if nums[left]+nums[right] > target {
			right--
		} else {
			res = (res + flag[right-left]) % mod
			left++
		}
	}
	return res
}
