package question_471_480

// 476. 数字的补数
// https://leetcode-cn.com/problems/number-complement
// Topics: 位运算

func findComplement(num int) int {
	var res = 0
	for i := 0; num > 0; num, i = num>>1, i+1 {
		res = (num&1^1)<<i + res
	}
	return res
}
