package question_251_260

// 258. 各位相加
// https://leetcode-cn.com/problems/add-digits

func addDigits(num int) int {
	t := num
	res := num
	for res > 9 {
		res = 0
		for t >= 1 {
			res += t % 10
			t /= 10
		}
		t = res
	}
	return res
}
