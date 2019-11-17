package question_51_60

// 53. 最大子序和
// https://leetcode-cn.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
		if dp[i] < nums[i] {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
