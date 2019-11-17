package question_221_230

// 224. 基本计算器
// https://leetcode-cn.com/problems/basic-calculator/
// Topics: 栈 数学

func cal(f uint8) func(a, b int) int {
	if f == '+' {
		return func(a, b int) int {
			return a + b
		}
	} else {
		return func(a, b int) int {
			return a - b
		}
	}
}

func calculate(s string) int {
	s = s + "+"
	res, n, fun := 0, 0, cal('+')
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else if s[i] == '+' || s[i] == '-' {
			res, fun, n = fun(res, n), cal(s[i]), 0
		} else if s[i] == '(' {
			k, t := i+1, 1
			for ; t > 0; i++ {
				if s[i+1] == '(' {
					t++
				} else if s[i+1] == ')' {
					t--
				}
			}
			n = calculate(s[k:i])
		}
	}
	return res
}
