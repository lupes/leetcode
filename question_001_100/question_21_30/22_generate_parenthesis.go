package question_0011_0020

import (
	"strings"
)

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
