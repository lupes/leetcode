package question_11_20

// 14. 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix
// Topics: 字符串

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longPrefix := strs[0]
	for _, str := range strs {
		if len(str) < len(longPrefix) {
			longPrefix = str
		}
	}
	for i := len(longPrefix); i >= 0; i-- {
		j := 0
		for ; j < len(strs) && strs[j][:i] == longPrefix[:i]; j++ {
		}
		if j == len(strs) {
			return longPrefix[:i]
		}
	}
	return ""
}
