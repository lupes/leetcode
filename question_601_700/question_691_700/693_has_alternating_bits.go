package question_691_700

// 693. 交替位二进制数
// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/

func hasAlternatingBits(n int) bool {
	if n&1 == 1 {
		n <<= 1
	}
	for n > 0 {
		if n&3 != 2 {
			return false
		}
		n >>= 2
	}
	return true
}
