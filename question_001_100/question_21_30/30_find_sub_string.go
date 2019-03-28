package question_0011_0020

import (
	"sort"
	"strings"
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
	wordMap := make(map[string]int)
	var values []string
	for _, v := range words {
		n, ok := wordMap[v]
		wordMap[v] = n + 1
		if !ok {
			values = append(values, v)
		}
	}

	wordCount := len(words)
	wordLen := len(words[0])
	tmp := make(map[int]string)
	var keys []int
	for _, value := range values {
		index := 0
		for {
			j := strings.Index(s[index:], value)
			if j > -1 {
				tmp[j+index] = value
				keys = append(keys, j+index)
				index = j + index + 1
			} else {
				break
			}
		}
	}
	var res []int
	sort.Ints(keys)
L:
	for _, key := range keys {
		w := tmp[key]
		t := make(map[string]int)
		t[w] = 1
		for i := 1; i < wordCount; i++ {
			j := key + i*wordLen
			v, ok := tmp[j]
			if !ok {
				continue L
			} else {
				t[v] = t[v] + 1
			}
		}
		for _, v := range values {
			if t[v] != wordMap[v] {
				continue L
			}
		}
		res = append(res, key)
	}
	return res
}
