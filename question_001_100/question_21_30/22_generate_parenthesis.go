package question_0011_0020

import (
	"strings"
)

// 22. 括号生成
// https://leetcode-cn.com/problems/generate-parentheses

func generateParenthesis(n int) []string {
	res := &[]string{}
	recursive(res, n, n, "")
	return *res
}

func recursive(res *[]string, n, m int, str string) {
	var i int
	if n == m {
		i = 1
	}
	for ; i < n+1; i++ {
		t := str + strings.Repeat("(", i)
		if n == i {
			*res = append(*res, t+strings.Repeat(")", m))
		} else {
			recursive(res, n-i, m-1, t+")")
		}
	}
}
