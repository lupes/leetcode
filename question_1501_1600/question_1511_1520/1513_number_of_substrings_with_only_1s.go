package question_1511_1520

// 1513. 仅含 1 的子串数
// https://leetcode-cn.com/problems/number-of-substrings-with-only-1s/
// Topics: 字符串 数组

func numSub(s string) int {
	s += "0"
	left, res, flag := 0, 0, false
	for i, c := range s {
		if !flag && c == '1' {
			left = i
			flag = true
		} else if flag && c == '0' {
			res += ((i - left + 1) * (i - left) / 2) % 1000000007
			flag = false
		}
	}
	return res
}
