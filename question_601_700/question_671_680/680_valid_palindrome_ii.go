package question_671_680

// 680. 验证回文字符串 Ⅱ
// https://leetcode-cn.com/problems/valid-palindrome-ii
// Topics: 字符串

func validPalindrome(s string) bool {
	var left, right = 0, len(s) - 1
	for right > left && s[left] == s[right] {
		left, right = left+1, right-1
	}
	if right == left || right == left+1 || right < left {
		return true
	}
	return isPalindrome(s[left+1:right+1]) || isPalindrome(s[left:right])
}

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; right > left; left, right = left+1, right-1 {
		if s[right] != s[left] {
			return false
		}
	}
	return true
}
