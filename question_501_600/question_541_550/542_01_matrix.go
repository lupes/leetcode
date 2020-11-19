package question_541_550

// 542. 01 矩阵
// https://leetcode-cn.com/problems/01-matrix/
// Topics: 深度优先搜索 广度优先搜索

var matrixFlag = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func updateMatrix(matrix [][]int) [][]int {
	r := len(matrix)
	if r < 0 {
		return matrix
	}
	c := len(matrix[0])
	queue, point := make([][2]int, 0), [2]int{0, 0}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			matrix[i][j] *= -1
			if matrix[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for len(queue) > 0 {
		point, queue = queue[0], queue[1:]
		for _, p := range matrixFlag {
			i, j := point[0]+p[0], point[1]+p[1]
			if i >= 0 && i < r && j >= 0 && j < c && matrix[i][j] == -1 {
				matrix[i][j] = matrix[point[0]][point[1]] + 1
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	return matrix
}
