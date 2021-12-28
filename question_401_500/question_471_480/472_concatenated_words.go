package question_471_480

import (
	"sort"
)

// 472. 连接词
// https://leetcode-cn.com/problems/concatenated-words
// Topics: 深度优先搜索 字典树 动态规划

type node struct {
	end   bool
	child map[int32]*node
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var root = &node{
		end:   false,
		child: make(map[int32]*node, 26),
	}

	var res []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		if findAllConcatenatedWordsInADictHelper(root, word) {
			res = append(res, word)
		} else {
			addWord(root, word)
		}
	}

	return res
}

func addWord(root *node, word string) {
	var next = root
	for _, c := range word {
		tmp := next.child[c-'a']
		if tmp == nil {
			tmp = &node{
				end:   false,
				child: make(map[int32]*node, 3),
			}
			next.child[c-'a'] = tmp
		}
		next = tmp
	}
	next.end = true
}

func findAllConcatenatedWordsInADictHelper(root *node, word string) bool {
	var dp = make([]bool, len(word)+1)
	dp[0] = true
	for i := range word {
		if !dp[i] {
			continue
		}
		var next = root
		for j := i; j < len(word); j++ {
			tmp := next.child[int32(word[j]-'a')]
			if tmp == nil {
				break
			}
			if tmp.end {
				dp[j+1] = true
			}
			next = tmp
		}
	}

	return dp[len(word)]
}
