package question_941_950

// 946. 验证栈序列
// https://leetcode-cn.com/problems/validate-stack-sequences
// Topics: 栈

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	for {
		if len(stack) == 0 && len(pushed) == 0 && len(popped) == 0 {
			return true
		} else if len(stack) > 0 && len(popped) > 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1]
			popped = popped[1:]
		} else if len(pushed) > 0 {
			stack = append(stack, pushed[0])
			pushed = pushed[1:]
		} else {
			return false
		}
	}
}
