package question_1361_1370

// 1380. 矩阵中的幸运数
// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
// Topics: 数组

func luckyNumbers(matrix [][]int) []int {
	var res []int
	var rl, cl = len(matrix), len(matrix[0])
	for i := 0; i < rl; i++ {
		var min, index = 100000, 0
		for j := 0; j < cl; j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				index = j
			}
		}
		var max = 0
		for k := 0; k < rl; k++ {
			if matrix[k][index] > max {
				max = matrix[k][index]
			}
		}
		if min == max {
			res = append(res, min)
		}
	}
	return res
}
