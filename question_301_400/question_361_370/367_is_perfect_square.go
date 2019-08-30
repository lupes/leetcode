package question_361_370

// 367. 有效的完全平方数
// https://leetcode-cn.com/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	i := 1
	for ; i*i < num; i++ {
	}
	return i*i == num
}
