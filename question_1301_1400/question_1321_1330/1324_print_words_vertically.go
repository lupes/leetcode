package question_1321_1330

import (
	"strings"
)

// 1324. 竖直打印单词
// https://leetcode-cn.com/problems/print-words-vertically/
// Topics: 字符串

func printVertically(s string) []string {
	arr := strings.Split(s, " ")
	var res []string
	for i := 0; ; i++ {
		a, n := "", 0
		for _, t := range arr {
			if i < len(t) {
				a += t[i : i+1]
			} else {
				a += " "
				n++
			}
		}
		if n < len(arr) {
			res = append(res, strings.TrimRight(a, " "))
		} else {
			break
		}
	}
	return res
}
