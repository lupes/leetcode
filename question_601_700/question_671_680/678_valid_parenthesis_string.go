package question_671_680

// 678. 有效的括号字符串
// https://leetcode-cn.com/problems/valid-parenthesis-string
// Topics: 字符串

func checkValidString(s string) bool {
	var stack []int32
	var left = 0
	for _, c := range s {
		switch c {
		case '*':
			stack = append(stack, c)
		case '(':
			left++
			stack = append(stack, c)
		case ')':
			if left > 0 {
				for i := len(stack) - 1; i >= 0; i-- {
					if stack[i] == '(' {
						stack = append(stack[:i], stack[i+1:]...)
						break
					}
				}
				left--
			} else if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	var tmp = 0
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == '*' {
			tmp++
		} else if stack[i] == '(' {
			if tmp > 0 {
				tmp--
			} else {
				return false
			}
		}
	}
	return true
}
