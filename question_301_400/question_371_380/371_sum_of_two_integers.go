package question_371_380

// 371. 两整数之和
// https://leetcode-cn.com/problems/sum-of-two-integers
// Topics: 位运算

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
