package question_171_180

// 172. 阶乘后的零
// https://leetcode-cn.com/problems/factorial-trailing-zeroes

func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res += n / 5
		n /= 5
	}
	return res
}
