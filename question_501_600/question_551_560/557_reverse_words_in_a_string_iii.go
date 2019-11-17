package question_551_560

// 557. 反转字符串中的单词 III
// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

func reverseWords(s string) string {
	var res, tmp = "", ""
	s += " "
	for i, c := range s {
		tmp = s[i:i+1] + tmp
		if c == ' ' {
			res, tmp = res+tmp, ""
		}
	}
	return res[1:]
}
