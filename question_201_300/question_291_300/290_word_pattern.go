package question_291_300

import "strings"

// 290. 单词规律
// https://leetcode-cn.com/problems/word-pattern/

func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	if len(pattern) != len(arr) {
		return false
	}
	m1, m2 := make(map[uint8]string), make(map[string]uint8)
	for i, s := range arr {
		v1, ok := m1[pattern[i]]
		if !ok {
			m1[pattern[i]] = s
		} else if v1 != s {
			return false
		}
		v2, ok := m2[s]
		if !ok {
			m2[s] = pattern[i]
		} else if v2 != pattern[i] {
			return false
		}
	}
	return true
}
