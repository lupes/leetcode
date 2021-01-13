package question_1331_1340

// 1347. 制造字母异位词的最小步骤数
// https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
// Topics: 字符串

func minSteps(s string, t string) int {
	var flag = [26]int{}
	for i := range s {
		flag[s[i]-'a']++
		flag[t[i]-'a']--
	}

	var res = 0
	for _, n := range flag {
		if n > 0 {
			res += n
		}
	}
	return res
}
