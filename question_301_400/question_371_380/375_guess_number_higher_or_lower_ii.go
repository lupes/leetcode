package question_371_380

// 375. 猜数字大小 II
// https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii
// Topics: 极小化极大 动态规划

func getMoneyAmount(n int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			f[i][j] = 40000
			for k := i; k < j; k++ {
				var minCost = k + f[k+1][j]
				if f[i][k-1] > f[k+1][j] {
					minCost = k + f[i][k-1]
				}
				if minCost < f[i][j] {
					f[i][j] = minCost
				}
			}
		}
	}
	return f[1][n]
}
