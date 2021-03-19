package question_1541_1550

// 1541. 平衡括号字符串的最少插入次数
// https://leetcode-cn.com/problems/minimum-insertions-to-balance-a-parentheses-string/
// Topics: 栈 字符串

func minInsertions(s string) int {
	var stack, res = make([]byte, 0), 0

	// 将字符串根据三种情况转换(:0 )):1 ):1 res+1
	for i := 0; i < len(s); {
		if s[i] == '(' {
			stack, i = append(stack, 0), i+1
		} else if i+2 <= len(s) && s[i:i+2] == "))" {
			stack, i = append(stack, 1), i+2
		} else {
			stack, res, i = append(stack, 1), res+1, i+1
		}
		if len(stack) > 1 && stack[len(stack)-2] == 0 && stack[len(stack)-1] == 1 {
			stack = stack[:len(stack)-2]
		}
	}

	for _, n := range stack {
		if n == 1 {
			res += 1
		} else {
			res += 2
		}
	}

	return res
}
