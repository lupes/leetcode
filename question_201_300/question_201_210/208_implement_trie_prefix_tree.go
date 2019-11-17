package question_201_210

import "strings"

// 208. 实现 Trie (前缀树)
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// Topics: 设计 字典树

type Trie struct {
	DataMap map[string]struct{}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		DataMap: make(map[string]struct{}, 0),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	this.DataMap[word] = struct{}{}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	_, ok := this.DataMap[word]
	return ok
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for key, _ := range this.DataMap {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
