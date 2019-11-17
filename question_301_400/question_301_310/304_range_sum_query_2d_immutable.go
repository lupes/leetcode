package question_301_310

// 304. 二维区域和检索 - 矩阵不可变
// https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

type NumMatrix struct {
	Matrix [][]int
}

// Topics: 动态规划

func ConstructorMatrix(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{
			Matrix: [][]int{{0}},
		}
	}
	if len(matrix[0]) == 0 {
		return NumMatrix{
			Matrix: [][]int{{0}},
		}
	}
	m := NumMatrix{}
	m.Matrix = make([][]int, len(matrix)+1)
	m.Matrix[0] = make([]int, len(matrix[0])+1)
	for r, nums := range matrix {
		m.Matrix[r+1] = make([]int, len(nums)+1)
		for c, n := range nums {
			m.Matrix[r+1][c+1] = n + m.Matrix[r+1][c] + m.Matrix[r][c+1] - m.Matrix[r][c]
		}
	}
	return m
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.Matrix[row2+1][col2+1] - this.Matrix[row2+1][col1] - this.Matrix[row1][col2+1] + this.Matrix[row1][col1]
}
