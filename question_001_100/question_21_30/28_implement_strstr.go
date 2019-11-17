package question_0011_0020

// 28. 实现 strStr()
// https://leetcode-cn.com/problems/implement-strstr

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hLen := len(haystack)
	nLen := len(needle)
	if nLen > hLen {
		return -1
	}
	for i, c := range haystack {
		if hLen-i < nLen {
			break
		}
		if uint8(c) == needle[0] {
			flag := true
			for j := 0; j < nLen; j++ {
				if haystack[i+j] != needle[j] {
					flag = false
				}
			}
			if flag {
				return i
			}
		}
	}
	return -1
}
