package question_1321_1330

import (
	"sort"
)

// 1329. 将矩阵按对角线排序
// https://leetcode-cn.com/problems/print-words-vertically/
// Topics: 排序 数组

func diagonalSort(mat [][]int) [][]int {
	rl, cl := len(mat), len(mat[0])
	for i := 0; i < cl-1; i++ {
		var tmp []int
		for j := 0; j < rl && i+j < cl; j++ {
			tmp = append(tmp, mat[j][i+j])
		}
		sort.Ints(tmp)
		for j := 0; j < rl && i+j < cl; j++ {
			mat[j][i+j] = tmp[j]
		}
	}
	for i := 1; i < rl-1; i++ {
		var tmp []int
		for j := 0; j < cl && i+j < rl; j++ {
			tmp = append(tmp, mat[i+j][j])
		}
		sort.Ints(tmp)
		for j := 0; j < cl && i+j < rl; j++ {
			mat[i+j][j] = tmp[j]
		}
	}
	return mat
}
