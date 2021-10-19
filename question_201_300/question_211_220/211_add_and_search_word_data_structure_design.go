package question_211_220

// 211. 添加与搜索单词 - 数据结构设计
// https://leetcode-cn.com/problems/add-and-search-word-data-structure-design
// Topics: 设计 字典树 回溯算法

type WordDictionary struct {
	words map[int][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		words: make(map[int][]string, 0),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.words[len(word)] = append(this.words[len(word)], word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	for _, tmp := range this.words[len(word)] {
		var res = true
		for i := range tmp {
			if !(tmp[i] == word[i] || tmp[i] == '.' || word[i] == '.') {
				res = false
				break
			}
		}
		if res {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
