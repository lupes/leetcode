package question_1611_1620

// 1614. 括号的最大嵌套深度
// https://leetcode-cn.com/problems/maximum-nesting-depth-of-the-parentheses/
// Topics: 栈 字符串

func maxDepth(s string) int {
	var stack, max int
	for _, c := range s {
		if c == '(' {
			stack++
		} else if c == ')' {
			stack--
		}
		if stack > max {
			max = stack
		}
	}
	return max
}
