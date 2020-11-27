package question_291_300

// 300. 最长上升子序列
// https://leetcode-cn.com/problems/longest-increasing-subsequence
// Topics: 二分查找 动态规划

func lengthOfLIS(nums []int) int {
	var dp, tmp, max = make([]int, len(nums)), 0, 0
	for i, n := range nums {
		tmp = 0
		for j := 0; j < i; j++ {
			if n > nums[j] && dp[j] > tmp {
				tmp = dp[j]
			}
		}
		dp[i] = tmp + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
