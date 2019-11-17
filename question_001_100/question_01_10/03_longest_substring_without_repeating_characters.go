package question_01_10

// 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// Topics: 哈希表 双指针 字符串

func lengthOfLongestSubstring(s string) int {
	var flag, max, t = make(map[byte]int), 0, 0
	for left, right := 0, 0; right < len(s); right++ {
		if n, ok := flag[s[right]]; ok && n+1 > left {
			t = right - left
			left = n + 1
		} else {
			t = right - left + 1
		}
		if t > max {
			max = t
		}
		flag[s[right]] = right
	}
	return max
}
