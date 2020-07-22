package question_531_540

// 537. 复数乘法
// https://leetcode-cn.com/problems/complex-number-multiplication
// Topics: 数学 字符串

func complexNumberMultiply(a string, b string) string {
	c, d := parseComplexNum(a)
	e, f := parseComplexNum(b)
	return toStr(c*e-d*f) + "+" + toStr(c*f+d*e) + "i"
}

func parseComplexNum(s string) (a, b int) {
	if s[len(s)-1] == 'i' {
		s = s[:len(s)-1]
	}
	for i, c := range s {
		if c == '+' {
			return toInt(s[:i]), toInt(s[i+1:])
		}
	}
	return toInt(s), 0
}

func toStr(i int) string {
	if i == 0 {
		return "0"
	}
	var res, f = "", ""
	if i < 0 {
		f, i = "-", i*-1
	}
	for i > 0 {
		res, i = string(i%10+'0')+res, i/10
	}
	return f + res
}

func toInt(s string) int {
	r, f := 0, 1
	if s[0] == '-' {
		s, f = s[1:], -1
	}
	for _, c := range s {
		r = r*10 + int(c-'0')
	}
	return r * f
}
