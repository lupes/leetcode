package question_541_550

// 541. 反转字符串 II
// https://leetcode-cn.com/problems/reverse-string-ii/
// Topics: 字符串

func reverseStr(s string, k int) string {
	l, k2 := len(s), 2*k
	var res string
	for i := 0; i < l; i += k2 {
		if i+k > l {
			res += reverse(s[i:])
		} else if i+k2 > l {
			res += reverse(s[i:i+k]) + s[i+k:]
		} else {
			res += reverse(s[i:i+k]) + s[i+k:i+k2]
		}
	}
	return res
}

func reverse(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		res = s[i:i+1] + res
	}
	return res
}
