package question_641_650

// 650. 只有两个键的键盘
// https://leetcode-cn.com/problems/2-keys-keyboard
// Topics: 动态规划

func minSteps(n int) int {
	var dp = make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 2; j*j <= i; j++ {
			if i%j == 0 && dp[i/j]+j < dp[i] {
				dp[i] = dp[i/j] + j
			}
		}
	}
	return dp[n]
}
