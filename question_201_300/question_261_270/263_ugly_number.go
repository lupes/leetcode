package question_261_270

// 263. 丑数
// https://leetcode-cn.com/problems/ugly-number/
// Topics: 数学

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num%5 == 0 {
		num /= 5
	}
	for num%3 == 0 {
		num /= 3
	}

	for num%2 == 0 {
		num /= 2
	}
	return num == 1
}
