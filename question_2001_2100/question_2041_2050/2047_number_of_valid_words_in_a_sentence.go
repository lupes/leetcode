package question_2041_2050

import (
	"fmt"
	"regexp"
	"strings"
)

// 2047. 句子中的有效单词数
// https://leetcode-cn.com/problems/number-of-valid-words-in-a-sentence/
// Topic: 字符串

func countValidWords(sentence string) int {
	var res = 0
	reg, _ := regexp.Compile(`(^[a-z]+(-[a-z]+)?[!,.]?$)|(^[!,.]$)`)
	arr := strings.Split(sentence, " ")
	for _, s := range arr {
		if reg.MatchString(s) {
			fmt.Printf("[%s]\n", s)
			res++
		}
	}

	return res
}
