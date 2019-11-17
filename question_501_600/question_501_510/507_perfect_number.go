package question_501_510

import "math"

// 507. 完美数
// https://leetcode-cn.com/problems/perfect-number/
// Topics: 数学

func checkPerfectNumber(num int) bool {
	var sum, t = 1, int(math.Sqrt(float64(num)))
	for i := 2; i <= t; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
	}
	if t*t == num {
		sum -= t
	}
	return sum == num && num > 1
}
