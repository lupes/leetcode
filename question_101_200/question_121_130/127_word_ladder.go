package question_121_130

// 127. 单词接龙
// https://leetcode-cn.com/problems/word-ladder
// Topics: 广度优先搜索

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var m = make(map[string]map[string]struct{})
	src := append(wordList, beginWord)
	for i := 0; i < len(src); i++ {
		for _, word := range wordList {
			t := 0
			for k := range word {
				if word[k] != src[i][k] {
					t++
					if t == 2 {
						break
					}
				}
			}
			if t == 1 {
				if _, ok := m[word]; !ok {
					m[word] = make(map[string]struct{})
				}
				m[word][src[i]] = struct{}{}
				if _, ok := m[src[i]]; !ok {
					m[src[i]] = make(map[string]struct{})
				}
				m[src[i]][word] = struct{}{}
			}
		}
	}
	if _, ok := m[endWord]; !ok {
		return 0
	}
	var res = 1
	var now = []string{beginWord}
	var next []string
	for len(now) > 0 {
		for _, word := range now {
			list, ok := m[word]
			if ok {
				for k, _ := range list {
					if k == endWord {
						return res + 1
					}
					next = append(next, k)
					if n, ok := m[k]; ok {
						delete(n, word)
						if len(n) == 0 {
							delete(m, k)
						}
					}
				}
			}
		}
		res++
		now = next
		next = make([]string, 0)
	}
	return 0
}
