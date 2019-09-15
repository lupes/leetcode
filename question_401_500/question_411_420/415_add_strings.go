package question_411_420

// 415. 字符串相加
// https://leetcode-cn.com/problems/add-strings/

func addStrings(num1 string, num2 string) string {
	l1, l2, n, c, res := len(num1), len(num2), uint8(0), uint8(0), ""
	for i := 0; i < l1 || i < l2 || c > 0; i++ {
		n, c = c, 0
		if i < l1 {
			n += num1[l1-i-1] - '0'
		}
		if i < l2 {
			n += num2[l2-i-1] - '0'
		}
		if n >= 10 {
			c = 1
			n -= 10
		}
		res = string(n+'0') + res
	}
	return res
}
