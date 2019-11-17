package question_361_370

// 367. 有效的完全平方数
// https://leetcode-cn.com/problems/valid-perfect-square/
// Topics: 数学 二分查找

func isPerfectSquare(num int) bool {
	i := 1
	for ; i*i < num; i++ {
	}
	return i*i == num
}
