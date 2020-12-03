package question_201_210

// 208. 实现 Trie (前缀树)
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// Topics: 设计 字典树

type Trie struct {
	End bool
	Son map[uint8]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		End: false,
		Son: make(map[uint8]*Trie, 26),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.End = true
		return
	}
	son, ok := this.Son[word[0]]
	if !ok {
		son = &Trie{
			Son: make(map[uint8]*Trie, 26),
		}
	}
	son.Insert(word[1:])
	this.Son[word[0]] = son
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.End
	}
	son, ok := this.Son[word[0]]
	if !ok {
		return false
	}
	return son.Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	c, last := prefix[0], prefix[1:]
	son, ok := this.Son[c]
	if !ok {
		return false
	}
	return son.StartsWith(last)
}
