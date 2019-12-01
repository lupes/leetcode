package question_131_140

// 139. 单词拆分
// https://leetcode-cn.com/problems/word-break
// Topics: 动态规划

func wordBreak(s string, wordDict []string) bool {
	var l = len(s)
	var res = make([]bool, l+1)
	res[0] = true
	for i := 1; i <= l; i++ {
		for _, word := range wordDict {
			wl := len(word)
			if i >= wl {
				if s[i-wl:i] == word && res[i-wl] {
					res[i] = true
					break
				}
			}
		}
	}
	return res[l]
}
