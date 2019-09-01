package question_381_390

// 387. 字符串中的第一个唯一字符
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	var m = make([]int32, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
