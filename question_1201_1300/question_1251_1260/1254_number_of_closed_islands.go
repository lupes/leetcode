package question_1251_1260

// 1254. 统计封闭岛屿的数目
// https://leetcode-cn.com/problems/number-of-closed-islands
// Topics: 深度优先搜索

var flag = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func closedIsland(grid [][]int) int {
	lr, lc := len(grid), len(grid[0])
	var res = 0
	for r, row := range grid {
		for c, n := range row {
			if n == 0 && closedIslandDFS(grid, lr, lc, r, c) {
				res++
			}
		}
	}
	return res
}

func closedIslandDFS(grid [][]int, lr, lc, r, c int) bool {
	if r < 0 || r >= lr || c < 0 || c >= lc {
		return false
	}
	if grid[r][c] >= 1 {
		return true
	}
	grid[r][c] = 2
	var res = true
	for _, f := range flag {
		res = closedIslandDFS(grid, lr, lc, r+f[0], c+f[1]) && res
	}
	return res
}
