package question_521_530

// 521. 最长特殊序列 Ⅰ
// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i
// Topics: 字符串

func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if a != b {
		return len(b)
	}
	return -1
}
