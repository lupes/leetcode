package question_341_350

// 344. 反转字符串
// https://leetcode-cn.com/problems/reverse-string/

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}
