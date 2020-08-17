package question_991_1000

// 991. 坏了的计算器
// https://leetcode-cn.com/problems/broken-calculator
// Topics: 贪心算法 数学

func brokenCalc(X int, Y int) int {
	var res = 0
	for ; Y > X; res++ {
		if Y%2 == 0 {
			Y /= 2
		} else {
			Y += 1
		}
	}
	return res + X - Y
}
