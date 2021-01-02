package question_1241_1250

// 1249. 移除无效的括号
// https://leetcode-cn.com/problems/minimum-remove-to-make-valid-parentheses
// Topics: 栈 字符串

func minRemoveToMakeValid(s string) string {
	var stack []int
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i+1)
		} else if c == ')' && (len(stack) == 0 || (len(stack) > 0 && stack[len(stack)-1] < 0)) {
			stack = append(stack, -(i + 1))
		} else if c == ')' && len(stack) > 0 && stack[len(stack)-1] > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return s
	}

	var res = make([]byte, 0, len(s)-len(stack))
	for i, c := range s {
		if len(stack) > 0 && (stack[0] == i+1 || stack[0] == -(i+1)) {
			stack = stack[1:]
		} else {
			res = append(res, byte(c))
		}
	}
	return string(res)
}
