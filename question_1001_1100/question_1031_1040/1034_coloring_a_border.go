package question_1031_1040

// 1034. 边框着色
// https://leetcode-cn.com/problems/coloring-a-border
// Topics: 深度优先搜索

var flag = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	rl, cl := len(grid), len(grid[0])
	src := grid[r0][c0]
	if src == color {
		return grid
	}

	var mask = make([][]bool, len(grid))
	for i := range grid {
		mask[i] = make([]bool, len(grid[i]))
	}

	colorBorderHelp(grid, mask, rl, cl, r0, c0, src, color)

	for i := range grid {
		for j := range grid[i] {
			if !mask[i][j] {
				continue
			}
			if !isBoundary(mask, rl, cl, i, j) {
				grid[i][j] = src
			}
		}
	}

	return grid
}

func isBoundary(mask [][]bool, rl, cl, i, j int) bool {
	if i == 0 || j == 0 || i == rl-1 || j == cl-1 {
		return true
	}
	for _, t := range flag {
		r, c := i+t[0], j+t[1]
		if !mask[r][c] {
			return true
		}
	}
	return false
}

func colorBorderHelp(grid [][]int, mask [][]bool, rl, cl, r, c, src, color int) {
	if r < 0 || r >= rl || c < 0 || c >= cl || grid[r][c] != src {
		return
	}
	grid[r][c] = color
	mask[r][c] = true
	for _, t := range flag {
		colorBorderHelp(grid, mask, rl, cl, r+t[0], c+t[1], src, color)
	}
}
