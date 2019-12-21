package question_221_230

// 227. 基本计算器 II
// https://leetcode-cn.com/problems/basic-calculator-ii
// Topics: 字符串

func calculate2(s string) int {
	s += "="
	var stack = make([]int32, 1)
	var flag = '0'
	for _, i := range s {
		if i >= '0' && i <= '9' {
			stack[len(stack)-1] = stack[len(stack)-1]*10 + (i - '0')
		} else if i == '+' || i == '-' || i == '*' || i == '/' || i == '=' {
			if flag == '*' {
				stack[len(stack)-3] *= stack[len(stack)-1]
				stack = stack[:len(stack)-2]
			} else if flag == '/' {
				stack[len(stack)-3] /= stack[len(stack)-1]
				stack = stack[:len(stack)-2]
			}
			stack = append(stack, i, 0)
			if i == '*' || i == '/' {
				flag = i
			} else {
				flag = '0'
			}
		}
	}
	var res = stack[0]
	for i := 2; i < len(stack); i += 2 {
		if stack[i-1] == '+' {
			res += stack[i]
		} else {
			res -= stack[i]
		}
	}
	return int(res)
}
