package question_341_350

// https://leetcode-cn.com/problems/integer-break/

// 343. 整数拆分
// https://leetcode-cn.com/problems/integer-break
// Topics: 数学 动态规划

func integerBreak(n int) int {
	var dp = make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			t := j * (i - j)
			if j*dp[i-j] > t {
				t = j * dp[i-j]
			}
			if t > dp[i] {
				dp[i] = t
			}
		}
	}
	return dp[n]
}
