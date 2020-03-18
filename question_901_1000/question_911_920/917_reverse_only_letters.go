package question_911_920

// 917. 仅仅反转字母
// https://leetcode-cn.com/problems/reverse-only-letters
// Topics: 字符串

func isLetter(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func reverseOnlyLetters(S string) string {
	left, right := 0, len(S)-1
	var res = make([]byte, len(S))
	for right >= left {
		if isLetter(S[left]) && isLetter(S[right]) {
			res[left], res[right] = S[right], S[left]
			left++
			right--
		} else if !isLetter(S[left]) {
			res[left] = S[left]
			left++
		} else if !isLetter(S[right]) {
			res[right] = S[right]
			right--
		}
	}
	return string(res)
}
