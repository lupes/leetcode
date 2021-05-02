package question_1661_1670

// 1662. 检查两个字符串数组是否相等
// https://leetcode-cn.com/problems/check-if-two-string-arrays-are-equivalent/
// Topics: 字符串

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1, s2 []byte
	for _, word := range word1 {
		s1 = append(s1, []byte(word)...)
	}

	for _, word := range word2 {
		s2 = append(s2, []byte(word)...)
	}

	return string(s1) == string(s2)
}
