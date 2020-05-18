package question_1001_1010

// 1003. 检查替换后的词是否有效
// https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions
// Topics: 栈 字符串

func isValid(S string) bool {
	var stack []byte
	for _, c := range S {
		c, l := byte(c), len(stack)
		if l > 1 && c == 'c' && stack[l-1] == 'b' && stack[l-2] == 'a' {
			stack = stack[:l-2]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
