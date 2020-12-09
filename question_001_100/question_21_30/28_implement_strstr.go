package question_0011_0020

// 28. 实现 strStr()
// https://leetcode-cn.com/problems/implement-strstr
// Topics: 双指针 字符串

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	l := len(needle)
	for i := 0; i <= len(haystack)-l; i++ {
		if haystack[i] == needle[0] && haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}
