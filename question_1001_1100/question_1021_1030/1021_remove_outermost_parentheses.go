package question_1021_1030

// 1021. 删除最外层的括号
// https://leetcode-cn.com/problems/remove-outermost-parentheses
// Topics: 栈

func removeOuterParentheses(S string) string {
	var stack, res, start = make([]byte, 0), "", 0
	for i, c := range S {
		if c == '(' {
			stack = append(stack, '(')
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res += S[start+1 : i]
				start = i + 1
			}
		}
	}
	return res
}
