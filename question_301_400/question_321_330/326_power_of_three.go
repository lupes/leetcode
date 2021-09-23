package question_321_330

// 326. 3的幂
// https://leetcode-cn.com/problems/power-of-three/
// Topics: 数学

func isPowerOfThree(n int) bool {
	switch n {
	case 1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441,
		1594323, 4782969, 14348907, 43046721, 129140163, 387420489, 1162261467:
		return true
	}
	return false

	// return n > 0 && 1162261467 % n == 0
}
