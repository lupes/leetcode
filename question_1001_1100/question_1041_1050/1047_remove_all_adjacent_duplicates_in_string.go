package question_1041_1050

// 1047. 删除字符串中的所有相邻重复项
// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
// Topics: 栈

func removeDuplicates(S string) string {
	var stack []byte
	for _, c := range S {
		c := byte(c)
		if len(stack) > 0 && c == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
