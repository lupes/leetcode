package question_761_770

// 766. 托普利茨矩阵
// https://leetcode-cn.com/problems/toeplitz-matrix
// Topics: 数组

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}
