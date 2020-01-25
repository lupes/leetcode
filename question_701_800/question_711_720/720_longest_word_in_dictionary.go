package question_711_720

import (
	"sort"
)

// 720. 词典中最长的单词
// https://leetcode-cn.com/problems/longest-word-in-dictionary
// Topics: 字典树 哈希表

type node struct {
	l   int
	son map[string]node
}

func longestWord(words []string) string {
	var res []string
	var max int
	var flag = make(map[string]struct{})
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	for _, word := range words {
		flag[word] = struct{}{}
	}
Next:
	for _, n := range words {
		l := len(n)
		if max != 0 && l < max {
			goto Over
		}
		for i := len(n) - 1; i > 0; i-- {
			if _, ok := flag[n[:i]]; !ok {
				continue Next
			}
		}
		res = append(res, n)
		max = l
	}
Over:
	sort.Strings(res)
	if len(res) == 0 {
		return ""
	}
	return res[0]
}
