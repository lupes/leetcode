package question_161_170

import (
	"strconv"
)

// 166. 分数到小数
// https://leetcode-cn.com/problems/fraction-to-recurring-decimal
// Topics: 哈希表 数学

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	if denominator == 0 {
		return ""
	}
	var m = make(map[int]int)
	var res string
	if numerator*denominator < 0 {
		res = "-"
	}
	if numerator < 0 {
		numerator *= -1
	}
	if denominator < 0 {
		denominator *= -1
	}
	for j := 0; numerator != 0; j++ {
		t := numerator / denominator
		if t <= -1 || t >= 1 {
			res += strconv.Itoa(t)
			numerator -= t * denominator
		} else {
			res += "0"
		}
		if j == 0 && numerator != 0 {
			res += "."
		}
		numerator *= 10
		if l, ok := m[numerator]; ok {
			return res[:l] + "(" + res[l:] + ")"
		} else {
			m[numerator] = len(res)
		}
	}
	return res
}
