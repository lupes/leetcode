package question_631_640

import (
	"math"
)

// 633. 平方数之和
// https://leetcode-cn.com/problems/sum-of-square-numbers
// Topics: 数学

func judgeSquareSum(c int) bool {
	t := math.Sqrt(float64(c))
	for i := 0; i <= int(t); i++ {
		j := int(math.Sqrt(float64(c - i*i)))
		if j*j+i*i == c {
			return true
		}
	}
	return false
}
