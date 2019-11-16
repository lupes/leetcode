package question_61_70

// 69. x 的平方根
// https://leetcode-cn.com/problems/sqrtx

func mySqrt(x int) int {
	l, r := 1, x
	for i := (l + r) / 2; i >= 0; i = (l + r) / 2 {
		if i*i == x {
			return i
		} else if i*i > x {
			r = i
		} else if (i+1)*(i+1) > x {
			return i
		} else {
			l = i
		}
	}
	return 0
}
