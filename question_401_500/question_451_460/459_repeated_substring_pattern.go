package question_451_460

// 459. 重复的子字符串
// https://leetcode-cn.com/problems/repeated-substring-pattern
// Topics: 字符串

func repeatedSubstringPattern(s string) bool {
	t := s + s
	for i := 1; i < len(s); i++ {
		if s == t[i:len(s)+i] {
			return true
		}
	}
	return false
}
