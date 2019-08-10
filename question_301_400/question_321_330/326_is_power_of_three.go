package question_321_330

// 326. 3的幂
// https://leetcode-cn.com/problems/power-of-three/

func isPowerOfThree(n int) bool {
	if n < 3 {
		return false
	}
	for ; n > 1; n /= 3 {
		if n%3 != 0 {
			return false
		}
	}
	return true
}
