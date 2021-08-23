package question_1641_1650

// 1646. 获取生成数组中的最大值
// https://leetcode-cn.com/problems/get-maximum-in-generated-array/
// Topics: 数组 动态规划 模拟

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	var dp, res = make([]int, n+1), 0
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + dp[i/2+1]
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
