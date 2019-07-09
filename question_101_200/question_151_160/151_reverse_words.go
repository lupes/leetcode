package question_151_160

import "strings"

// 151. 翻转字符串里的单词
// https://leetcode-cn.com/problems/reverse-words-in-a-string/

func reverseWords(s string) string {
	tmp := strings.Split(s, " ")
	var arr []string
	for _, s := range tmp {
		if s != "" {
			arr = append([]string{s}, arr...)
		}
	}
	return strings.Join(arr, " ")
}
