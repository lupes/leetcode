package question_1321_1330

// 1328. 破坏回文串
// https://leetcode-cn.com/problems/print-words-vertically/
// Topics: 字符串

func breakPalindrome(palindrome string) string {
	if len(palindrome) <= 1 {
		return ""
	}
	tmp := []byte(palindrome)
	for i := 0; i < len(tmp)/2; i++ {
		if tmp[i] != 'a' {
			tmp[i] = 'a'
			return string(tmp)
		}
	}
	tmp[len(tmp)-1] = 'b'
	return string(tmp)
}
