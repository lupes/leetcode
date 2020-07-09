package question_391_400

// 400. 第N个数字
// https://leetcode-cn.com/problems/nth-digit
// Topics: 数学

var flag400 = [][2]int{{0, 0}, {9, 9}, {189, 99}, {2889, 999}, {38889, 9999}, {488889, 99999}, {5888889, 999999},
	{68888889, 9999999}, {788888889, 99999999}, {8888888889, 999999999}, {98888888889, 9999999999}}

func findNthDigit(n int) int {
	for i, t := range flag400 {
		if n < t[0] {
			a, b := (n-flag400[i-1][0])/i, (n-flag400[i-1][0])%i
			if b == 0 {
				return (a + flag400[i-1][1]) % 10
			} else {
				c := (a + flag400[i-1][1]) + 1
				for j := 0; j < (i - b); j++ {
					c /= 10
				}
				return c % 10
			}
		}
	}
	return 0
}
