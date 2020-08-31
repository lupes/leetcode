package question_1181_1190

// 1190. 反转每对括号间的子串
// https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses
// Topics: 栈

func reverseParentheses(s string) string {
	var stack []byte
	var flag []int
	for _, c := range s {
		if c == ')' {
			for start, end := flag[len(flag)-1], len(stack)-1; end > start; start, end = start+1, end-1 {
				stack[start], stack[end] = stack[end], stack[start]
			}
			flag = flag[:len(flag)-1]
		} else if c == '(' {
			flag = append(flag, len(stack))
		} else {
			stack = append(stack, byte(c))
		}
	}
	return string(stack)
}
