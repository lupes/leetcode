package question_1451_1460

// 1455. 检查单词是否为句中其他单词的前缀
// https://leetcode-cn.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
// Topics: 字符串

func isPrefixOfWord(sentence string, searchWord string) int {
	var res, tmp, lw, ls = 1, " " + sentence, len(searchWord), len(sentence) + 1
	for i := 0; i < ls-lw; i++ {
		if tmp[i] == ' ' {
			if lw+i+1 <= ls && searchWord == tmp[i+1:i+lw+1] {
				return res
			} else {
				res++
			}
		}
	}
	return -1
}
