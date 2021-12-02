package question_1441_1450

// 1446. 连续字符
// https://leetcode-cn.com/problems/consecutive-characters/
// Topics: 字符串

func maxPower(s string) int {
	var res, start = 0, 0
	for i := range s {
		if s[start] == s[i] {
			if i-start+1 > res {
				res = i - start + 1
			}
		} else {
			start = i
		}
	}
	return res
}
