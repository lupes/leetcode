package question_01_10

// 3. 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// Topics: 哈希表 双指针 字符串

func lengthOfLongestSubstring(s string) int {
	var flag = make(map[byte]int)
	var left, right, max, t = 0, 0, 0, true
	for right < len(s) {
		if t {
			c := s[right]
			if n, ok := flag[c]; ok && n == 1 {
				flag[c]++
				t = false
			} else {
				flag[c]++
				if right-left+1 > max {
					max = right - left + 1
				}
			}
			right++
		} else {
			c := s[left]
			if n, ok := flag[c]; ok && n == 2 {
				flag[c]--
				t = true
			} else {
				flag[c]--
			}
			left++
		}
	}

	return max
}
