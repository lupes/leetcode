package question_111_120

// 115. 不同的子序列
// https://leetcode-cn.com/problems/distinct-subsequences/

func numDistinct(s string, t string) int {
	sl, tl := len(s), len(t)
	dp := make([][]int, tl+1)
	dp[0] = make([]int, sl+1)
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i, tc := range t {
		dp[i+1] = make([]int, sl+1)
		for j, sc := range s {
			if sc == tc {
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}
	return dp[tl][sl]
}
