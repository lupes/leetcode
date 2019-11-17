package question_491_500

// 504. 七进制数
// https://leetcode-cn.com/problems/base-7/
// Topics:

func convertToBase7(num int) string {
	res, flag := "", true
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num, flag = num*-1, false
	}
	for num != 0 {
		res = string(byte(num%7)+'0') + res
		num /= 7
	}
	if !flag {
		return "-" + res
	}
	return res
}
