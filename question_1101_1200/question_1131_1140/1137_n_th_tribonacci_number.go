package question_1131_1140

// 1137. 第 N 个泰波那契数
// https://leetcode-cn.com/problems/n-th-tribonacci-number
// Topics: 递归

func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	if n == 0 {
		return a
	} else if n == 1 {
		return b
	} else if n == 2 {
		return c
	}
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
