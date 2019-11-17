package question_01_10

// 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	length1, length2 := 0, 0
	for i := 0; i < len(s); i++ {
		var old string
	Break:
		for j := i; j < len(s); j++ {
			old = s[i : j+1]
			for k := 0; k < len(old)-1; k++ {
				if s[j] == old[k] {
					length2 = len(old) - 1
					break Break
				}
			}
			length2 = len(old)
		}
		if length2 > length1 {
			length1 = length2
			length2 = 0
		}
	}
	return length1
}
