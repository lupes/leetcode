package question_1371_1380

// 1400. 构造 K 个回文字符串
// https://leetcode-cn.com/problems/construct-k-palindrome-strings/
// Topics: 贪心算法

func canConstruct(s string, k int) bool {
	l := len(s)
	if l < k {
		return false
	} else if l == k {
		return true
	} else if l >= 26 && k >= 26 {
		return true
	}

	var flag = [26]int{}
	for _, c := range s {
		flag[c-'a']++
	}

	var res = 0
	for _, n := range flag {
		if n%2 == 1 {
			res += 1
		}
	}
	return res <= k
}
