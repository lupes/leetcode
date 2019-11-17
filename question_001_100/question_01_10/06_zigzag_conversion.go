package question_01_10

import (
	"strings"
)

// 6. Z 字形变换
// https://leetcode-cn.com/problems/zigzag-conversion
// Topics: 字符串

func convert(s string, numRows int) string {
	if s == "" {
		return ""
	}
	if numRows == 1 {
		return s
	}
	var tmp [][]string
	for i := 0; i < numRows; i++ {
		tmp = append(tmp, []string{})
	}
	for i := range s {
		a := i / (numRows - 1)
		b := i % (numRows - 1)
		if a%2 == 0 {
			tmp[b] = append(tmp[b], s[i:i+1])
		} else {
			tmp[numRows-b-1] = append(tmp[numRows-b-1], s[i:i+1])
		}
	}
	res := ""
	for _, a := range tmp {
		res += strings.Join(a, "")
	}
	return res
}
