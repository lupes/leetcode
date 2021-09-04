package offer

// 剑指 Offer 10- I. 斐波那契数列
// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// Topics: 记忆化搜索 数学 动态规划

func fib(n int) int {
	if n < 2 {
		return n
	}
	var fibNums = make([]int, n+1)
	fibNums[0], fibNums[1] = 0, 1
	for i := 2; i <= n; i++ {
		fibNums[i] = (fibNums[i-1] + fibNums[i-2]) % (1e9 + 7)
	}
	return fibNums[n]
}
