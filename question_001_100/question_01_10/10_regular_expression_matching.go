package question_01_10

// 10. 正则表达式匹配
// https://leetcode-cn.com/problems/regular-expression-matching
// Topics: 字符串 动态规划 回溯算法

func isMatch(s string, p string) bool {
	var m []string
	for _, c := range p {
		if c == '*' {
			m[len(m)-1] += "*"
		} else {
			m = append(m, string(c))
		}
	}
	var dp = make([][]bool, len(m)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true
	for i, t := range m {
		if len(t) == 2 && dp[i][0] {
			dp[i+1][0] = true
		}
		for j := range s {
			if len(t) == 2 && dp[i][j] {
				dp[i+1][j] = true
			}
			if (len(t) == 2 && dp[i+1][j] && (t[0] == s[j] || t[0] == '.')) ||
				(len(t) == 2 && dp[i][j+1]) ||
				(t[0] == '.' && dp[i][j]) ||
				(t[0] == s[j] && dp[i][j]) {
				dp[i+1][j+1] = true
			}
		}
	}
	return dp[len(m)][len(s)]
}
