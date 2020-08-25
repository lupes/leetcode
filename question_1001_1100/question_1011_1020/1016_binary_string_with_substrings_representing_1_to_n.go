package question_1011_1020

import (
	"strconv"
	"strings"
)

// 1016. 子串能表示从 1 到 N 数字的二进制串
// https://leetcode-cn.com/problems/binary-string-with-substrings-representing-1-to-n
// Topics: 字符串

func queryString(S string, N int) bool {
	for i := int64(1); i <= int64(N); i++ {
		if !strings.Contains(S, strconv.FormatInt(i, 2)) {
			return false
		}
	}
	return true
}
