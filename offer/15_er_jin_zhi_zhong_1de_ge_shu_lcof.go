package offer

// 剑指 Offer 15. 二进制中1的个数
// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
// Topics: 位运算

func hammingWeight(num uint32) int {
	var res = 0
	for ; num > 0; num >>= 1 {
		res += int(num & 1)
	}
	return res
}
