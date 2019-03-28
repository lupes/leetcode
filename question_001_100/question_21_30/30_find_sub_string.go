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
	wordsLen := wordCount * wordLen
	tmp := make([]string, wordCount, wordCount)
	for i := 0; i < size-wordsLen+1; i++ {
		for j := 0; j < wordCount; j++ {
			start := i + j*wordLen
			tmp[j] = s[start : start+wordLen]
		}
		sort.Strings(tmp)
		if reflect.DeepEqual(tmp, words) {
			res = append(res, i)
		}
	}
	return res
}
