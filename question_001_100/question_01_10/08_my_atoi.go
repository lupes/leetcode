package question_01_10

import (
	"math"
)

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	res := 0
	flag := 1
	start := true
	for _, c := range str {
		if start {
			if c == ' ' {
				continue
			} else if c == '-' {
				flag = -1
			} else if c == '+' {
				flag = 1
			} else if c >= '0' && c <= '9' {
				res = int(c - '0')
			} else {
				break
			}
			start = false
		} else {
			if c >= '0' && c <= '9' {
				res = res*10 + int(c-'0')
				if res*flag > math.MaxInt32 {
					return math.MaxInt32
				}
				if res*flag < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		}
	}
	return res * flag
}
