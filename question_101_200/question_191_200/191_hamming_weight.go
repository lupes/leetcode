package question_191_200

// 191. 位1的个数
// https://leetcode-cn.com/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	var res int
	for num > 0 {
		res += int(num & 1)
		num >>= 1
	}
	return res
}
