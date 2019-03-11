package question_11_20

import "strings"

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	res := ""
	m := num / 1000
	c := num % 1000 / 100
	x := num % 1000 % 100 / 10
	i := num % 1000 % 100 % 10
	res += helper(m, "M", "", "")
	res += helper(c, "C", "D", "M")
	res += helper(x, "X", "L", "C")
	res += helper(i, "I", "V", "X")
	return res
}

func helper(a int, one, five, ten string) string {
	if a == 9 {
		return one + ten
	} else if 8 >= a && a >= 5 {
		return five + strings.Repeat(one, a-5)
	} else if a == 4 {
		return one + five
	} else {
		return strings.Repeat(one, a)
	}
}
