package question_211_220

// 214. 最短回文串
// https://leetcode-cn.com/problems/shortest-palindrome
// Topics: 字符串

func shortestPalindrome(s string) string {
	l, i := len(s), 0
	if l == 0 || l == 1 {
		return s
	}
	for i = l; i > 1 && !isPalindrome(s[:i]); i-- {
	}
	res := s
	for ; i < l; i++ {
		res = s[i:i+1] + res
	}
	return res
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	l, r := len(s)/2-1, len(s)/2+len(s)%2
	for ; l >= 0; l, r = l-1, r+1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
