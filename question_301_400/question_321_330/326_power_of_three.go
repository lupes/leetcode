package question_321_330

// 326. 3的幂
// https://leetcode-cn.com/problems/power-of-three/
// Topics: 数学

func isPowerOfThree(n int) bool {
	return n > 0 && 4052555153018976267%n == 0
}
