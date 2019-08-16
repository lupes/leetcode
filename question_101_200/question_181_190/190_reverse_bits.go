package question_181_190

// 190. 颠倒二进制位
// https://leetcode-cn.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	var res, l uint32
	for num > 0 {
		res <<= 1
		res += num & 1
		num >>= 1
		l++
	}
	return res << (32 - l)
}
