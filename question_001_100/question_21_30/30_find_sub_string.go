package question_0011_0020

import (
	"reflect"
	"sort"
)

func findSubstring(s string, words []string) []int {
	if s == "" {
		return nil
	}
	if len(words) == 0 {
		return nil
	}
	if words[0] == "" {
		return nil
	}
	var res []int
	sort.Strings(words)
	size := len(s)
	wordCount := len(words)
	wordLen := len(words[0])
	tmp := make([]string, wordCount, wordCount)
	for i := 0; i < wordLen; i++ {
		var ss []string
		for j := 0; i+(j+1)*wordLen <= size; j++ {
			ss = append(ss, s[i+j*wordLen:i+(j+1)*wordLen])
		}
		for j := 0; j <= len(ss)-wordCount; j++ {
			copy(tmp, ss[j:j+wordCount])
			sort.Strings(tmp)
			if reflect.DeepEqual(tmp, words) {
				res = append(res, i+j*wordLen)
			}
		}
	}
	return res
}
