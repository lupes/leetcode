package question_121_130

// 125. 验证回文串
// https://leetcode-cn.com/problems/valid-palindrome

func isValid(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for r > l {
		vr, vl := isValid(s[r]), isValid(s[l])
		if vr && vl {
			if s[r] == s[l] || (s[r] >= 'a' && s[r]-32 == s[l]) || (s[l] >= 'a' && s[l]-32 == s[r]) {
				r--
				l++
			} else {
				return false
			}
		} else {
			if !vr {
				r--
			}
			if !vl {
				l++
			}
		}
	}
	return true
}
