package question_51_60

// 58. 最后一个单词的长度
// https://leetcode-cn.com/problems/length-of-last-word
// Topics: 字符串

func lengthOfLastWord(s string) int {
	var last, tmp = 0, 0
	for _, c := range s {
		if c == ' ' {
			tmp = 0
		} else {
			tmp += 1
			last = tmp
		}
	}
	return last
}
