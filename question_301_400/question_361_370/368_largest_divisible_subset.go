package question_361_370

import (
	"sort"
)

// 368. 最大整除子集
// https://leetcode-cn.com/problems/largest-divisible-subset
// Topics: 数学 动态规划

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	var max, index, dp = 1, 0, make([][2]int, len(nums))
	for i, n := range nums {
		dp[i] = [2]int{1, -1}
		for j := i - 1; j >= 0; j-- {
			if n%nums[j] == 0 && dp[j][0]+1 > dp[i][0] {
				dp[i][0], dp[i][1] = dp[j][0]+1, j
			}
		}
		if dp[i][0] > max {
			max, index = dp[i][0], i
		}
	}

	var res []int
	for index > -1 {
		res = append(res, nums[index])
		index = dp[index][1]
	}

	return res
}
