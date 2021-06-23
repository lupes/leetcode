package question_311_320

// 316. 去除重复字母
// https://leetcode-cn.com/problems/remove-duplicate-letters
// Topics: 栈 贪心算法

func removeDuplicateLetters(s string) string {
	var flag, tmp ['z' + 1]int
	var stack []byte
	for _, c := range s {
		flag[c]++
	}
	for _, c := range s {
		flag[c]--
		for len(stack) > 0 && flag[stack[len(stack)-1]] > 0 && byte(c) < stack[len(stack)-1] && tmp[c] == 0 {
			tmp[stack[len(stack)-1]]--
			stack = stack[:len(stack)-1]
		}
		if tmp[c] == 0 {
			stack = append(stack, byte(c))
			tmp[c]++
		}
	}
	return string(stack)
}
