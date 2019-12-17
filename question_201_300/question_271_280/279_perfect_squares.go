package question_271_280

import (
	"math"
)

// 279. 完全平方数
// https://leetcode-cn.com/problems/perfect-squares
// Topics: 广度优先搜索 数学 动态规划

func numSquares(n int) int {
	var res, t = make([]int, n+1), 0
	for i := 1; i <= n; i++ {
		t = i
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if res[i-j*j]+1 < t {
				t = res[i-j*j] + 1
			}
		}
		res[i] = t
	}
	return res[n]
}
