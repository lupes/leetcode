package question_01_10

// 5. 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	o, n := s[:1], ""
	for i := range s {
		for j, k := i-1, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				n = s[j : k+1]
			} else {
				if len(n) > len(o) {
					o = n
				}
				break
			}
		}
		for j, k := i, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				n = s[j : k+1]
			} else {
				if len(n) > len(o) {
					o = n
				}
				break
			}
		}
	}
	return o
}
