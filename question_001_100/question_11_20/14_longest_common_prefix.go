package question_11_20

// 14. 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix
// Topics: 字符串

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		size := len(longPrefix)
		if len(strs[i]) < size {
			size = len(strs[i])
			longPrefix = longPrefix[:size]
		}
		for j := 0; j < size; j++ {
			if longPrefix[j] != strs[i][j] {
				longPrefix = longPrefix[0:j]
				if longPrefix == "" {
					return ""
				}
				break
			}
		}
	}
	return longPrefix
}
