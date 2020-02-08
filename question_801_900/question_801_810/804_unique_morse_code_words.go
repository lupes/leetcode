package question_801_810

import (
	"bytes"
)

// 804. 唯一摩尔斯密码词
// https://leetcode-cn.com/problems/unique-morse-code-words
// Topics: 字符串

var flag = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
	".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
	"...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	var tmp = make(map[string]struct{})
	for _, word := range words {
		var buffer = bytes.NewBufferString("")
		for _, c := range word {
			buffer.WriteString(flag[c-'a'])
		}
		tmp[buffer.String()] = struct{}{}
	}
	return len(tmp)
}
