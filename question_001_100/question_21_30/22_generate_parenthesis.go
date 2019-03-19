package question_0011_0020

import (
	"strings"
)

func generateParenthesis(n int) []string {
	return recursive(n, n, "")
}

func recursive(n, m int, str string) []string {
	var i int
	if n == m {
		i = 1
	}
	var res []string
	for ; i < n+1; i++ {
		t := str + strings.Repeat("(", i)
		if n == i {
			res = append(res, t+strings.Repeat(")", m))
		} else {
			res = append(res, recursive(n-i, m-1, t+")")...)
		}
	}
	return res
}
