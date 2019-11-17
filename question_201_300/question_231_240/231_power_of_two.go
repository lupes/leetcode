package question_231_240

// 231. 2的幂
// https://leetcode-cn.com/problems/power-of-two/
// Topics: 位运算 数学

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}
