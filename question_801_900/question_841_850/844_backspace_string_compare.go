package question_841_850

// 844. 比较含退格的字符串
// https://leetcode-cn.com/problems/backspace-string-compare
// Topics: 栈 双指针

func backspaceCompare(S string, T string) bool {
	var a, b string
	for _, c := range S {
		if c == '#' && len(a) > 0 {
			a = a[:len(a)-1]
		} else if c != '#' {
			a += string(c)
		}
	}
	for _, c := range T {
		if c == '#' && len(b) > 0 {
			b = b[:len(b)-1]
		} else if c != '#' {
			b += string(c)
		}
	}
	return a == b
}
