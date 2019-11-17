package question_341_350

// 344. 反转字符串
// https://leetcode-cn.com/problems/reverse-string/
// Topics: 双指针 字符串

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
