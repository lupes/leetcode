package question_141_150

import (
	"strconv"
)

// 150. 逆波兰表达式求值
// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation
// Topics: 栈

func evalRPN(tokens []string) int {
	var arr []int
	var l int
	for _, s := range tokens {
		if s == "+" {
			arr[l-2] = arr[l-2] + arr[l-1]
		} else if s == "-" {
			arr[l-2] = arr[l-2] - arr[l-1]
		} else if s == "*" {
			arr[l-2] = arr[l-2] * arr[l-1]
		} else if s == "/" {
			arr[l-2] = arr[l-2] / arr[l-1]
		} else {
			n, _ := strconv.Atoi(s)
			arr = append(arr, n)
			l++
			continue
		}
		arr = arr[:l-1]
		l--
	}
	return arr[0]
}
