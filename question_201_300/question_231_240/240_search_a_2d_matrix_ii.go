package question_231_240

// 240. 搜索二维矩阵 II
// https://leetcode-cn.com/problems/search-a-2d-matrix-ii
// Topics: 二分查找 分治算法

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	r, c, cl := len(matrix)-1, 0, len(matrix[0])
	for r >= 0 && c < cl {
		t := matrix[r][c]
		if t == target {
			return true
		} else if t > target {
			r--
		} else if t < target {
			c++
		}
	}
	return false
}
