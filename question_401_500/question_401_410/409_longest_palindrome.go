package question_401_410

// 409. 最长回文串
// https://leetcode-cn.com/problems/longest-palindrome
// Topics: 哈希表

func longestPalindrome(s string) int {
	var flag = make(map[int32]int, 0)
	for _, c := range s {
		flag[c]++
	}
	res, t := 0, 0
	for _, v := range flag {
		if v%2 == 1 {
			t = 1
			res += v - 1
		} else {
			res += v
		}
	}
	if len(s) == 0 {
		return 0
	} else if res == 0 && len(s) > 0 {
		return 1
	} else if t == 0 {
		return res
	} else {
		return res + 1
	}
}
