package question_851_860

// 856. 括号的分数
// https://leetcode-cn.com/problems/score-of-parentheses
// Topics: 栈 字符串

func scoreOfParentheses(S string) int {
	var stack []int
	for _, c := range S {
		if c == '(' {
			stack = append(stack, -1)
		} else {
			last := -1
			for i := len(stack) - 1; i >= 0; i-- {
				if last == -1 && stack[i] == -1 {
					last = 1
					stack[i] = 1
				} else if last == -1 && stack[i] > 0 {
					last = stack[i] * 2
					stack[i-1] = last
					stack = stack[:i]
					i--
				} else if last > 0 && stack[i] > 0 {
					last = stack[i] + last
					stack[i] = last
					stack = stack[:i+1]
				} else if stack[i] == -1 {
					break
				}
			}
		}
	}
	return stack[0]
}
