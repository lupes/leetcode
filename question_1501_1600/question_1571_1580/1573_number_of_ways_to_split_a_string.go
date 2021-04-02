package question_1571_1580

// 1573. 分割字符串的方案数
// https://leetcode-cn.com/problems/number-of-ways-to-split-a-string/
// Topics: 字符串

func numWays(s string) int {
	num := 0
	for _, c := range s {
		if c == '1' {
			num++
		}
	}
	if len(s) < 3 {
		return 0
	} else if num == 0 {
		return (len(s) - 2) * (len(s) - 1) / 2 % (1e9 + 7)
	} else if num%3 != 0 {
		return 0
	}

	var left, right, n = 0, 0, 0
	for _, c := range s {
		if c == '1' {
			n++
		}
		if num/3 == n {
			left++
		}
		if num/3*2 == n {
			right++
		}
	}
	return left * right % (1e9 + 7)
}
