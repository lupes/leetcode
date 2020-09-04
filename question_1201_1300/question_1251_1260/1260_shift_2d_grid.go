package question_1251_1260

// 1260. 二维网格迁移
// https://leetcode-cn.com/problems/shift-2d-grid/
// Topics: 数组

func shiftGrid(grid [][]int, k int) [][]int {
	var rl, cl = len(grid), len(grid[0])
	k %= rl * cl
	if k == 0 {
		return grid
	}
	var res = make([][]int, rl)
	for i := range res {
		res[i] = make([]int, cl)
	}
	for i, row := range grid {
		for j, n := range row {
			t := (i*cl + j + k) % (rl * cl)
			res[t/cl][t%cl] = n
		}
	}
	return res
}
