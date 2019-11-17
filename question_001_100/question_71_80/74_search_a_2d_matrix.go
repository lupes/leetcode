package question_71_80

// 74. 搜索二维矩阵
// https://leetcode-cn.com/problems/search-a-2d-matrix

func searchMatrix(matrix [][]int, target int) bool {
	ll := len(matrix)
	if ll == 0 {
		return false
	}
	lc := len(matrix[0])
	if lc == 0 {
		return false
	}
	for _, row := range matrix {
		if row[0] > target {
			return false
		}
		if row[lc-1] >= target {
			for _, n := range row {
				if n == target {
					return true
				}
			}
			return false
		}
	}
	return false
}
