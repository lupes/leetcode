package question_921_930

// 921. 使括号有效的最少添加
// https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid
// Topics: 栈 贪心算法

func minAddToMakeValid(S string) int {
	var stack []int32
	for _, c := range S {
		if c == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack)
}
