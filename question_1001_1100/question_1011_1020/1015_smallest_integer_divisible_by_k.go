package question_1011_1020

// 1015. 可被 K 整除的最小整数
// https://leetcode-cn.com/problems/smallest-integer-divisible-by-k
// Topics: 数学

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	var i, t = 1, 1
	for ; t%K != 0; i, t = i+1, t*10+1 {
		t %= K
	}
	return i
}
