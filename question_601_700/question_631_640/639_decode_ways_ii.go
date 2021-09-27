package question_631_640

// 639. 解码方法 2
// https://leetcode-cn.com/problems/decode-ways-ii
// Topics: 动态规划

func numDecodings(s string) int {
	var dp = make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	if s[0] == '0' {
		dp[1] = 0
	} else if s[0] == '*' {
		dp[1] = 9
	}
	for i := 1; i < len(s); i++ {
		c1, c2 := s[i-1], s[i]
		a, b := 0, 0
		if (c1 == '0' || c1 > '2') && c2 == '0' {
			return 0
		} else if c1 == '*' && c2 == '*' {
			a, b = 15, 9
		} else if c1 == '*' && c2 == '0' {
			a, b = 2, 0
		} else if c1 == '*' && c2 > '6' {
			a, b = 1, 1
		} else if c1 == '*' && c2 <= '6' {
			a, b = 2, 1
		} else if c2 == '*' {
			b = 9
			for j := byte(1); c1 != '0' && j < 10 && (c1-'0')*10+j <= 26; j++ {
				a++
			}
		} else if c2 == '0' {
			a = 1
		} else {
			if c1 != '0' && (c1-'0')*10+c2-'0' <= 26 {
				a = 1
			}
			b = 1
		}
		dp[i+1] = (dp[i-1]*a + dp[i]*b) % (1e9 + 7)
	}
	return dp[len(s)]
}
