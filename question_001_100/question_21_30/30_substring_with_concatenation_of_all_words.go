package question_0011_0020

// 30. 串联所有单词的子串
// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
// Topics: 哈希表 双指针 字符串

func findSubstringHelper(wordMap map[string]int, s string, l int) bool {
	if len(s) == 0 && len(wordMap) == 0 {
		return true
	}
	var res bool
	n, ok := wordMap[s[:l]]
	if ok && n > 0 {
		wordMap[s[:l]]--
		if n == 1 {
			delete(wordMap, s[:l])
		}
		res = findSubstringHelper(wordMap, s[l:], l)
		wordMap[s[:l]]++
	}
	return res
}

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 || words[0] == "" {
		return nil
	}

	wordMap := make(map[string]int, len(words))
	for _, word := range words {
		wordMap[word]++
	}

	wordLen, wordsLen := len(words[0]), len(words)

	var res []int
	for i := 0; i <= len(s)-wordLen*wordsLen; i++ {
		if _, ok := wordMap[s[i:i+wordLen]]; ok {
			flag := findSubstringHelper(wordMap, s[i:i+wordLen*wordsLen], wordLen)
			if flag {
				res = append(res, i)
			}
		}
	}

	return res
}
