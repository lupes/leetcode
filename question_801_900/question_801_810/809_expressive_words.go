package question_801_810

import (
	"fmt"
)

// 809. 情感丰富的文字
// https://leetcode-cn.com/problems/expressive-words
// Topics: 字符串

func expressiveWords(S string, words []string) int {
	var flag = []int{1}
	var s = " "
	for i := range S {
		if S[i] == s[len(s)-1] {
			flag[len(flag)-1]++
		} else {
			flag = append(flag, 1)
			s += S[i : i+1]
		}
	}
	fmt.Printf("%s %+v\n", s, flag)
	var res = 0
	for _, word := range words {
		var tmp = " "
		var f = []int{1}
		for i := range word {
			if word[i] == tmp[len(tmp)-1] {
				f[len(f)-1]++
			} else {
				if tmp != s[:len(tmp)] || f[len(f)-1] > flag[len(f)-1] || !(flag[len(f)-1] >= 3 || flag[len(f)-1] == f[len(f)-1] || flag[len(f)-1]-f[len(f)-1] >= 2) {
					goto Next
				}
				f = append(f, 1)
				tmp += word[i : i+1]
			}
		}
		fmt.Printf("%s %s %+v\n", word, tmp, f)
		if tmp == s && f[len(f)-1] <= flag[len(flag)-1] && (flag[len(f)-1] >= 3 || flag[len(f)-1] == f[len(f)-1] || flag[len(f)-1]-f[len(f)-1] >= 2) {
			res++
		}
	Next:
	}
	return res
}
