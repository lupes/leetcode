package question_1071_1080

import (
	"strings"
)

// 1078. Bigram 分词
// https://leetcode-cn.com/problems/occurrences-after-bigram
// Topics: 哈希表

func findOcurrences(text string, first string, second string) []string {
	var res []string
	tmp := strings.Split(text, " ")
	for i := 2; i < len(tmp); i++ {
		if tmp[i-2] == first && tmp[i-1] == second {
			res = append(res, tmp[i])
		}
	}
	return res
}
