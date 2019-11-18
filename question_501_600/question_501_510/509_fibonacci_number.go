package question_501_510

// 509. 斐波那契数
// https://leetcode-cn.com/problems/fibonacci-number
// Topics: 数组

func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 || N == 2 {
		return 1
	} else {
		return fib(N-1) + fib(N-2)
	}
}
