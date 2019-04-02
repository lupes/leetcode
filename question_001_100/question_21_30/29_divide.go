package question_0011_0020

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	res := 0
	flag := dividend^divisor < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	for i := byte(31); ; i-- {
		if (dividend >> i) >= divisor {
			res += 1 << i
			dividend -= divisor << i
		}
		if dividend < divisor {
			break
		}
		if i == 0 {
			break
		}
	}
	if flag {
		res = -res
	}
	return res
}
