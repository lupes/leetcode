package question_1131_1140

// 1137. 第 N 个泰波那契数
// https://leetcode-cn.com/problems/n-th-tribonacci-number
// Topics: 递归

var flag = make([]int, 41)

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		flag[n] = 1
	} else if n == 2 {
		flag[n] = 1
	} else if flag[n] == 0 {
		flag[n] = tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
	}
	return flag[n]
}
