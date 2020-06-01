package question_1151_1160

// 1160. 拼写单词
// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters
// Topics: 数组 哈希表

func countCharacters(words []string, chars string) int {
	var flag = make([]int, 26)
	for _, c := range chars {
		flag[c-'a']++
	}
	var res = 0
	for _, word := range words {
		var tmp = make([]int, 26)
		for _, c := range word {
			tmp[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if flag[i] < tmp[i] {
				goto Next
			}
		}
		res += len(word)
	Next:
	}
	return res
}
