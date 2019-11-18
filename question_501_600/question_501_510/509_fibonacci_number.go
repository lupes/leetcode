package question_501_510

// 509. 斐波那契数
// https://leetcode-cn.com/problems/fibonacci-number
// Topics: 数组 动态规划

func fib(N int) int {
	var fibs = make([]int, N+2)
	fibs[1] = 1
	for i := 2; i <= N; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[N]
}
