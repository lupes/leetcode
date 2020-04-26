package question_1001_1010

// 1006. 笨阶乘
// https://leetcode-cn.com/problems/clumsy-factorial
// Topics: 数学

func clumsy(N int) int {
	switch N {
	case 3:
		return 6
	case 2:
		return 2
	case 1:
		return 1
	}
	a, b, c, d, t, res := 0, 0, 0, 0, N, 0
	if t >= 4 {
		a, b, c, d, t = t, t-1, t-2, t-3, t-4
		res = a*b/c + d
	}
	for t >= 4 {
		a, b, c, d, t = t, t-1, t-2, t-3, t-4
		res -= a*b/c - d
	}
	if t == 3 {
		res -= 6
	} else if t == 2 {
		res -= 2
	} else if t == 1 {
		res -= 1
	}
	return res
}
