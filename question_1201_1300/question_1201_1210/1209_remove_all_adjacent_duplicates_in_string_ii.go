package question_1201_1210

import (
	"strings"
)

// 1209. 删除字符串中的所有相邻重复项 II
// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string-ii
// Topics: 栈

func removeDuplicates(s string, k int) string {
	var stack [][2]int32
	for _, c := range s {
		l := len(stack)
		if l == 0 || (l > 0 && stack[l-1][0] != c) {
			stack = append(stack, [2]int32{c, 1})
		} else {
			stack[l-1][1]++
			if stack[l-1][1] == int32(k) {
				stack = stack[:l-1]
			}
		}
	}
	var res string
	for _, v := range stack {
		res += strings.Repeat(string(v[0]), int(v[1]))
	}
	return res
}
