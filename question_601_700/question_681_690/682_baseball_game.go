package question_681_690

import (
	"strconv"
)

// 682. 棒球比赛
// https://leetcode-cn.com/problems/baseball-game
// Topics: 栈

func calPoints(ops []string) int {
	var stack []int
	var res, l int
	for _, o := range ops {
		if o == "C" {
			res, stack, l = res-stack[l-1], stack[:l-1], l-1
			continue
		} else if o == "D" {
			stack = append(stack, stack[l-1]*2)
		} else if o == "+" {
			stack = append(stack, stack[l-2]+stack[l-1])
		} else {
			n, _ := strconv.Atoi(o)
			stack = append(stack, n)
		}
		res, l = res+stack[l], l+1
	}
	return res
}
