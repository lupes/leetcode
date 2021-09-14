package question_521_530

import (
	"sort"
)

// 524. 通过删除字母匹配到字典里最长单词
// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting
// Topics: 排序 双指针

func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		return len(d[i]) > len(d[j]) || (len(d[i]) == len(d[j]) && d[i] < d[j])
	})

	var sl = len(s)
	for _, word := range d {
		wl := len(word)
		if wl >= sl && word != s {
			continue
		}
		if wl == sl && word == s {
			return word
		}
		if findLongestWordHelper(s, word, sl, wl) {
			return word
		}
	}
	return ""
}

func findLongestWordHelper(s, w string, sl, wl int) bool {
	var si, wi = 0, 0
	for wi < wl && si < sl && (sl-si) >= (wl-wi) {
		if w[wi] == s[si] {
			wi++
		}
		si++
	}
	return wi == wl
}
