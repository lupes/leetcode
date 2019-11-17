package question_71_80

// 73. 矩阵置零
// https://leetcode-cn.com/problems/set-matrix-zeroes
// Topics: 数组

func setZeroes(matrix [][]int) {
	ll := len(matrix)
	if ll == 0 {
		return
	}
	lc := len(matrix[0])
	var flag = make([][]bool, ll)
	for i := range flag {
		flag[i] = make([]bool, lc)
	}
	for i, line := range matrix {
		for j, v := range line {
			if v == 0 {
				for k := 0; k < ll; k++ {
					flag[k][j] = true
				}
				for k := 0; k < lc; k++ {
					flag[i][k] = true
				}
			}
		}
	}
	for i, line := range matrix {
		for j := range line {
			if flag[i][j] {
				matrix[i][j] = 0
			}
		}
	}
}
