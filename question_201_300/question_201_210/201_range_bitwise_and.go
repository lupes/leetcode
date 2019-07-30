package question_201_210

import "math/bits"

// 201. 数字范围按位与
// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/

func rangeBitwiseAnd(m int, n int) int {
	if bits.Len(uint(m)) != bits.Len(uint(n)) {
		return 0
	}
	res := m
	for i := m + 1; i <= n; i++ {
		res = res & i
	}
	return res
}
