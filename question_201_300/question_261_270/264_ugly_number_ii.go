package question_261_270

// 264. 丑数 II
// https://leetcode-cn.com/problems/ugly-number-ii
// Topics: 堆 数学 动态规划

func nthUglyNumber(n int) int {
	var dp = make([]int, 0, n)
	dp = append(dp, 1)
	var i2, i3, i5 = 0, 0, 0
	for len(dp) < n {
		s2, s3, s5 := dp[i2]*2, dp[i3]*3, dp[i5]*5
		if s2 <= s3 && s2 <= s5 {
			if s2 > dp[len(dp)-1] {
				dp = append(dp, s2)
			}
			i2++
		} else if s3 <= s2 && s3 <= s5 {
			if s3 > dp[len(dp)-1] {
				dp = append(dp, s3)
			}
			i3++
		} else {
			if s5 > dp[len(dp)-1] {
				dp = append(dp, s5)
			}
			i5++
		}
	}
	return dp[n-1]
}
