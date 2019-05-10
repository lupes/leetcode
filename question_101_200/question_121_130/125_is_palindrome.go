package question_121_130

func isPalindrome(s string) bool {
	str := ""
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			str += string(c)
		} else if c >= 'A' && c <= 'Z' {
			str += string(c - 'A' + 'a')
		}
	}
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}
