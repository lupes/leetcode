package question_231_240

import (
	"math"
)

// 233. 数字 1 的个数
// https://leetcode-cn.com/problems/number-of-digit-one
// Topics: 数学

var m = []int{0, 1, 20, 300, 4000, 50000, 600000, 7000000, 80000000, 900000000}

func countDigitOne(n int) int {
	res, t := 0, 0
	for log := int(math.Log10(float64(n))); n > 0; log-- {
		mod := int(math.Pow10(log))
		t, n = n/mod, n%mod
		for j := 0; j < t; j++ {
			if j == 1 {
				res += mod
			}
			res += m[log]
		}
		if t == 1 {
			res += n + 1
		}
	}

	return res
}
