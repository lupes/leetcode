package question_311_320

import (
	"math"
)

// 313. 超级丑数
// https://leetcode-cn.com/problems/super-ugly-number
// Topics: 堆 数学

func nthSuperUglyNumber(n int, primes []int) int {
	var dp = make([]int, 0, n)
	dp = append(dp, 1)
	var indexes = make([]int, len(primes))
	for len(dp) < n {
		var min, i = math.MaxInt64, 0
		for j, prime := range primes {
			if prime*dp[indexes[j]] < min {
				min, i = prime*dp[indexes[j]], j
			}
		}
		if min > dp[len(dp)-1] {
			dp = append(dp, min)
		}
		indexes[i]++
	}
	return dp[n-1]
}
