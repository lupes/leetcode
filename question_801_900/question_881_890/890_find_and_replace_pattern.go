package question_881_890

// 890. 查找和替换模式
// https://leetcode-cn.com/problems/find-and-replace-pattern
// Topics: 字符串

func findAndReplacePattern(words []string, pattern string) []string {
	var res []string
Next:
	for _, word := range words {
		var wordFlag = make(map[byte]byte, 26)
		var patternFlag = make(map[byte]byte, 26)
		for i := range word {
			wordChar, patternOk := patternFlag[pattern[i]]
			patternChar, wordOk := wordFlag[word[i]]
			if !patternOk && !wordOk {
				patternFlag[pattern[i]] = word[i]
				wordFlag[word[i]] = pattern[i]
			} else if !(patternOk && wordOk && wordChar == word[i] && patternChar == pattern[i]) {
				continue Next
			}
		}
		res = append(res, word)
	}
	return res
}
