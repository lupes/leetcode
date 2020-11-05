package question_1361_1370

import (
	"math"
)

// 1362. 最接近的因数
// https://leetcode-cn.com/problems/closest-divisors/
// Topics: 数学

func closestDivisors(num int) []int {
	a := int(math.Sqrt(float64(num))) + 1
	b := a
	for {
		if a*b == num+1 || a*b == num+2 {
			return []int{a, b}
		} else if a*b < num+1 {
			b++
		} else {
			a--
		}
	}
}
