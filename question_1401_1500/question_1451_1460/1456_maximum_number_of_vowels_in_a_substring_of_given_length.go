package question_1451_1460

// 1456. 定长子串中元音的最大数目
// https://leetcode-cn.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
// Topics: 字符串 滑动窗口

func isVowels(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func maxVowels(s string, k int) int {
	max, tmp := 0, 0
	for i := range s {
		if isVowels(s[i]) {
			tmp++
		}
		if i >= k && isVowels(s[i-k]) {
			tmp--
		}
		if tmp > max {
			max = tmp
		}
	}

	return max
}
