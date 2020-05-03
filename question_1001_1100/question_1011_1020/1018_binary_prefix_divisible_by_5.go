package question_1011_1020

// 1018. 可被 5 整除的二进制前缀
// https://leetcode-cn.com/problems/binary-prefix-divisible-by-5
// Topics: 数组

func prefixesDivBy5(A []int) []bool {
	var res, t = make([]bool, len(A)), 0
	for i, n := range A {
		t <<= 1
		t += n
		t %= 10
		if t%5 == 0 {
			res[i] = true
		}
	}
	return res
}
