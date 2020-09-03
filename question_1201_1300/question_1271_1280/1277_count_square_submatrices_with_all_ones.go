package question_1261_1270

// 1277. 统计全为 1 的正方形子矩阵
// https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/
// Topics: 数组 动态规划

func countSquares(matrix [][]int) int {
	var res = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				continue
			}
			a, b, c := 0, 0, 0
			if i > 0 && j > 0 {
				a = matrix[i-1][j-1]
			}
			if i > 0 {
				b = matrix[i-1][j]
			}
			if j > 0 {
				c = matrix[i][j-1]
			}
			if b < a {
				a = b
			}
			if c < a {
				a = c
			}
			a++
			matrix[i][j] = a
			res += a
		}
	}
	return res
}
