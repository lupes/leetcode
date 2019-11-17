package question_01_10

// 7. 整数反转
// https://leetcode-cn.com/problems/reverse-integer
// Topics: 数学

func reverse(x int) int {
	res := 0
	if x == 0 {
		return 0
	} else if x > 0 {
		res = fun(x)
	} else {
		res = -fun(-x)
	}
	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res
}

func fun(x int) (res int) {
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	return res
}
