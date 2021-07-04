package question_591_600

import (
	"fmt"
)

// 592. 分数加减运算
// https://leetcode-cn.com/problems/fraction-addition-and-subtraction
// Topics: 数学

func fractionAddition(expression string) string {
	var flag = [11]int{0, 2520, 1260, 840, 630, 504, 420, 360, 315, 280, 252}
	var son, parent = 0, 2520
	for len(expression) > 0 {
		ts, tp := 1, 1
		if expression[0] == '-' {
			ts, expression = -1, expression[1:]
		} else if expression[0] == '+' {
			expression = expression[1:]
		}
		if expression[1] != '/' {
			ts, expression = ts*10, expression[3:]
		} else {
			ts, expression = ts*int(expression[0]-'0'), expression[2:]
		}
		if len(expression) > 1 && expression[1] != '+' && expression[1] != '-' {
			tp, expression = 10, expression[2:]
		} else {
			tp, expression = int(expression[0]-'0'), expression[1:]
		}
		son += ts * flag[tp]
	}
	if son == 0 {
		return "0/1"
	}
	m := mc(son, parent)
	son, parent = son/m, parent/m
	return fmt.Sprintf("%d/%d", son, parent)
}

func mc(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
