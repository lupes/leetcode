package question_131_140

// 131. 分割回文串
// https://leetcode-cn.com/problems/palindrome-partitioning
// Topics: 字符串 动态规划 回溯算法

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{{}}
	}
	var res [][]string
	for i := len(s) - 1; i >= 0; i-- {
		if isPalindrome(s[i:]) {
			for _, arr := range partition(s[:i]) {
				res = append(res, append(arr, s[i:]))
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	l := len(s)
	c := l / 2
	for i := 0; i <= c; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
