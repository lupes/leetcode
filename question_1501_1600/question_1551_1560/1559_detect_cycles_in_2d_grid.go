package question_1551_1560

// 1559. 二维网格图中探测环
// https://leetcode-cn.com/problems/detect-cycles-in-2d-grid/
// Topics: 深度优先搜索 广度优先搜索 并查集 数组 矩阵

var flag = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func containsCycle(grid [][]byte) bool {
	rl, cl := len(grid), len(grid[0])
	if rl == 1 || cl == 1 {
		return false
	}
	for i, row := range grid {
		for j, c := range row {
			if c >= 'a' && c <= 'z' {
				if findCycle(grid, rl, cl, -1, -1, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func findCycle(grid [][]byte, rl, cl, i1, j1, i2, j2 int) bool {
	grid[i2][j2] -= 32
	var t bool
	for _, f := range flag {
		a, b := i2+f[0], j2+f[1]
		if a < 0 || a >= rl || b < 0 || b >= cl || (a == i1 && b == j1) {
			continue
		}

		if grid[a][b] == grid[i2][j2] {
			return true
		}
		if grid[a][b]-32 == grid[i2][j2] {
			t = t || findCycle(grid, rl, cl, i2, j2, a, b)
		}
	}
	return t
}
