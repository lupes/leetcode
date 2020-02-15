package question_831_840

import (
	"sort"
)

// 833. 字符串中的查找与替换
// https://leetcode-cn.com/problems/find-and-replace-in-string
// Topics: 字符串

type replace struct {
	I int
	E int
	S string
	T string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	var replaces = make([]replace, len(indexes))
	for i, index := range indexes {
		replaces[i] = replace{
			I: index,
			E: index + len(sources[i]),
			S: sources[i],
			T: targets[i],
		}
	}
	sort.Slice(replaces, func(i, j int) bool {
		return replaces[i].I < replaces[j].I
	})
	var res = ""
	var start = 0
	for _, replace := range replaces {
		res += S[start:replace.I]
		if replace.E <= len(S) && S[replace.I:replace.E] == replace.S {
			res += replace.T
			start = replace.E
		} else {
			start = replace.I
		}
	}
	if start < len(S) {
		res += S[start:]
	}
	return res
}
