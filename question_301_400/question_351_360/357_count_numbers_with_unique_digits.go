package question_351_360

// 357. 计算各个位数不同的数字个数
// https://leetcode-cn.com/problems/count-numbers-with-unique-digits
// Topics: 数学 动态规划 回溯算法

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	}
	var sum, same = 10, 9
	for i := 0; i < n-1; i++ {
		same *= 9 - i
		sum += same
	}
	return sum
}
