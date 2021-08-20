package question_81_90

// 87. 扰乱字符串
// https://leetcode-cn.com/problems/scramble-string/
// Topics: 字符串 动态规划

var flag = make(map[string]bool, 100)

func isScrambleHelper(s1, s2 string) bool {
	t, ok := flag[s1+s2]
	if !ok {
		t = isScramble(s1, s2)
		flag[s1+s2] = t
	}
	return t
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) || (len(s1) == 1 && s1 != s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	for i := 1; i < len(s1); i++ {
		if isScrambleHelper(s1[:i], s2[:i]) && isScrambleHelper(s1[i:], s2[i:]) {
			return true
		}

		//if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
		//	return true
		//}

		if isScrambleHelper(s1[:i], s2[len(s1)-i:]) && isScrambleHelper(s1[i:], s2[:len(s1)-i]) {
			return true
		}

		//if isScramble(s1[:i], s2[len(s1)-i:]) && isScramble(s1[i:], s2[:len(s1)-i]) {
		//	return true
		//}
	}
	return false
}
