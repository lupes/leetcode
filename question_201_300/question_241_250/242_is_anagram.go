package question_241_250

// 242. 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sm = make(map[int32]int, 26)
	var tm = make(map[int32]int, 26)
	for _, i := range s {
		sm[i-'a']++
	}
	for _, i := range t {
		tm[i-'a']++
	}
	for i := int32(0); i < 26; i++ {
		if sm[i] != tm[i] {
			return false
		}
	}
	return true
}
