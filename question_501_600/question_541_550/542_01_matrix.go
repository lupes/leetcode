package question_541_550

// 542. 01 矩阵
// https://leetcode-cn.com/problems/01-matrix/
// Topics: 深度优先搜索 广度优先搜索

func updateMatrix(matrix [][]int) [][]int {
	r := len(matrix)
	if r < 0 {
		return matrix
	}
	c := len(matrix[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 1 {
				matrix[i][j] = -1
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == -1 {
				l, _ := updateMatrixDfs(matrix, r, c, i, j)
				matrix[i][j] = l
			}
		}
	}
	return matrix
}

func updateMatrixDfs(matrix [][]int, r, c, i, j int) (int, bool) {
	if (i == r || i == -1) || (j == c || j == -1) {
		return 0, false
	}
	if matrix[i][j] == -2 {
		return 0, false
	}
	if matrix[i][j] >= 0 {
		return matrix[i][j], true
	}
	matrix[i][j] = -2
	l := r + c
	if t, ok := updateMatrixDfs(matrix, r, c, i, j-1); ok && t < l {
		l = t
	}
	if t, ok := updateMatrixDfs(matrix, r, c, i-1, j); ok && t < l {
		l = t
	}
	if t, ok := updateMatrixDfs(matrix, r, c, i, j+1); ok && t < l {
		l = t
	}
	if t, ok := updateMatrixDfs(matrix, r, c, i+1, j); ok && t < l {
		l = t
	}
	matrix[i][j] = -1
	return l + 1, true
}
