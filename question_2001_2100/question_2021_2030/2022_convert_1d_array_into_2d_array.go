package question_2021_2030

// 2022. 将一维数组转变成二维数组
// https://leetcode-cn.com/problems/convert-1d-array-into-2d-array/
// Topic: 数组 矩阵 模拟

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != n*m {
		return nil
	}
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = original[i*n : i*n+n]
	}
	return matrix
}
