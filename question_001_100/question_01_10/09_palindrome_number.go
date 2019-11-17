package question_01_10

// 9. 回文数
// https://leetcode-cn.com/problems/palindrome-number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}
	res, tmp := 0, x
	for tmp > 0 {
		res = res*10 + tmp%10
		tmp = tmp / 10

		if tmp == res || tmp/10 == res {
			return true
		}
	}
	return res == x
}
