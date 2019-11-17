package question_341_350

import (
	"math"
)

// https://leetcode-cn.com/problems/integer-break/

// 343. 整数拆分
// https://leetcode-cn.com/problems/integer-break
// Topics: 数学 动态规划

func integerBreak(n int) int {
	m := int(math.Sqrt(float64(n)))
	if m == 1 {
		return m * (n - m)
	}
	t := m
	d := n - m*m
	if d > m {
		d = n - m*(m+1)
		m = m + 1
	}
	return int(math.Pow(float64(m), float64(t-d)) * math.Pow(float64(m+1), float64(d)))
}
