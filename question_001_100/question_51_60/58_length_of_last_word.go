package question_51_60

import "strings"

// 58. 最后一个单词的长度
// https://leetcode-cn.com/problems/length-of-last-word
// Topics: 字符串

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	s = strings.TrimRight(s, " ")
	size := len(s)
	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return len(s[i+1:])
		}
	}
	return size
}
