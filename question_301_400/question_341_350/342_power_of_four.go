package question_341_350

// 342. 4的幂
// https://leetcode-cn.com/problems/power-of-four/
// Topics: 位运算

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0xaaaaaaaaa == 0
}
