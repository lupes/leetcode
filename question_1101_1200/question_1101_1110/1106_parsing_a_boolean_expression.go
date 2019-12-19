package question_1101_1110

// 1106. 解析布尔表达式
// https://leetcode-cn.com/problems/parsing-a-boolean-expression
// Topics: 字符串

func parseBoolExpr(expression string) bool {
	var stack []byte
	for _, c := range expression {
		if c == ')' {
			j := len(stack) - 1
			var t1, t2, t3 byte = 't', 'f', 't'
			for ; j >= 0; j-- {
				if stack[j] == '(' {
					break
				} else if stack[j] == 't' {
					t1 = 'f'
					t2 = 't'
				} else if stack[j] == 'f' {
					t3 = 'f'
				}
			}
			switch stack[j-1] {
			case '!':
				stack[j-1] = t1
			case '|':
				stack[j-1] = t2
			case '&':
				stack[j-1] = t3
			}
			stack = stack[:j]
		} else if c != ',' {
			stack = append(stack, byte(c))
		}
	}
	return stack[0] == 't'
}
