package question_681_690

import (
	"strings"
)

// 686. 重复叠加字符串匹配
// https://leetcode-cn.com/problems/repeated-string-match
// Topics: 字符串

func repeatedStringMatch(A string, B string) int {
	var i, tmp = 1, A
	for ; len(tmp) < len(B); i++ {
		tmp += A
	}
	if strings.Contains(tmp, B) {
		return i
	}
	if strings.Contains(tmp+A, B) {
		return i + 1
	}
	return -1
}
