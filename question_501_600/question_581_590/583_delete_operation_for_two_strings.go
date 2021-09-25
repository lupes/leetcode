package question_581_590

// 583. 两个字符串的删除操作
// https://leetcode-cn.com/problems/delete-operation-for-two-strings
// Topics: 字符串

func minDistance(word1 string, word2 string) int {
	var dp = make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := range word1 {
		for j := range word2 {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = dp[i][j] + 2
				if dp[i][j+1]+1 < dp[i+1][j+1] {
					dp[i+1][j+1] = dp[i][j+1] + 1
				}
				if dp[i+1][j]+1 < dp[i+1][j+1] {
					dp[i+1][j+1] = dp[i+1][j] + 1
				}
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
