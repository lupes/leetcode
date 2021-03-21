package question_1541_1550

// 1544. 整理字符串
// https://leetcode-cn.com/problems/make-the-string-great/
// Topics: 栈 字符串

func makeGood(s string) string {
	var stack []byte
	for _, c := range s {
		if len(stack) > 0 && (byte(c-'A'+'a') == stack[len(stack)-1] || byte(c+'A'-'a') == stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(c))
		}
	}
	return string(stack)
}
