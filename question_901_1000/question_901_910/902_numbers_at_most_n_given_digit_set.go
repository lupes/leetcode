package question_901_910

// 902. 最大为 N 的数字组合
// https://leetcode-cn.com/problems/numbers-at-most-n-given-digit-set
// Topics: 数学 动态规划

func atMostNGivenDigitSet(D []string, N int) int {
	l, t, s := 0, N, ""
	for t > 0 {
		l, s, t = l+1, string(t%10+'0')+s, t/10
	}
	var dp = make([]int, l+1)
	dp[l] = 1
	for i := l - 1; i >= 0; i-- {
		for _, d := range D {
			if d[0] < s[i] {
				dp[i] += pow(len(D), l-i-1)
			} else if d[0] == s[i] {
				dp[i] += dp[i+1]
			}
		}
	}
	for i := 1; i < l; i++ {
		dp[0] += pow(len(D), i)
	}
	return dp[0]
}

func pow(x, y int) int {
	var res = 1
	for j := 0; j < y; j++ {
		res *= x
	}
	return res
}
