package question_1281_1290

// 1281. 整数的各位积和之差
// https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
// Topics: 数学

func subtractProductAndSum(n int) int {
	var a, b = 1, 0
	for n > 0 {
		t := n % 10
		a *= t
		b += t
		n /= 10
	}
	return a - b
}
