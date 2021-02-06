package question_1441_1450

import (
	"fmt"
)

// 1447. 最简分数
// https://leetcode-cn.com/problems/simplified-fractions/
// Topics: 数学

func simplifiedFractions(n int) []string {
	var res []string
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if !(i%2 == 0 && j%2 == 0) &&
				!(i%3 == 0 && j%3 == 0) &&
				!(i%5 == 0 && j%5 == 0) &&
				!(i%7 == 0 && j%7 == 0) &&
				simplifiedFractionsGCD(j, i) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return res
}

func simplifiedFractionsGCD(min, max int) int {
	t := max % min
	if t == 0 {
		return min
	}
	return simplifiedFractionsGCD(t, min)
}
