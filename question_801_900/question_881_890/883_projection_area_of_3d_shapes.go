package question_881_890

// 883. 三维形体投影面积
// https://leetcode-cn.com/problems/projection-area-of-3d-shapes
// Topics: 数学

func projectionArea(grid [][]int) int {
	var res, rows, cols = 0, make([]int, len(grid)), make([]int, len(grid))
	for i, row := range grid {
		for j, c := range row {
			if c > rows[i] {
				res = res - rows[i] + c
				rows[i] = c
			}
			if c > cols[j] {
				res = res - cols[j] + c
				cols[j] = c
			}
			if c > 0 {
				res += 1
			}
		}
	}
	return res
}
