package question_641_650

// 647. 回文子串
// https://leetcode-cn.com/problems/palindromic-substrings
// Topics: 字符串 动态规划

func countSubstrings(s string) int {
	var res, l = 0, len(s)
	for i := 0; i < l; i++ {
		res++
		for j := 1; i+j < l && i-j >= 0 && s[i-j] == s[i+j]; j++ {
			res++
		}
		for j := 0; i+j+1 < l && i-j >= 0 && s[i-j] == s[i+j+1]; j++ {
			res++
		}
	}
	return res
}
