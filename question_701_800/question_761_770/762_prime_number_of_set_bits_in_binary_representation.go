package question_761_770

import (
	"math/bits"
)

// 762. 二进制表示中质数个计算置位
// https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation
// Topics: 位运算

func countPrimeSetBits(L int, R int) int {
	var res = 0
	for i := L; i <= R; i++ {
		n := bits.OnesCount32(uint32(i))
		if n == 2 || n == 3 || n == 5 || n == 7 || n == 11 || n == 13 || n == 17 || n == 19 {
			res++
		}
	}
	return res
}
