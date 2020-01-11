package question_691_700

// 695. 岛屿的最大面积
// https://leetcode-cn.com/problems/max-area-of-island
// Topics: 深度优先搜索 数组

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var max, height, weight = 0, len(grid), len(grid[0])
	for i, row := range grid {
		for j, g := range row {
			if g == 1 {
				tmp := maxAreaOfIslandHelper(grid, height, weight, i, j)
				if tmp > max {
					max = tmp
				}
			}
		}
	}
	return max
}

func maxAreaOfIslandHelper(grid [][]int, height, weight, i, j int) int {
	if i < 0 || j < 0 || i >= height || j >= weight {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 2
	var res = 1
	res += maxAreaOfIslandHelper(grid, height, weight, i+1, j)
	res += maxAreaOfIslandHelper(grid, height, weight, i-1, j)
	res += maxAreaOfIslandHelper(grid, height, weight, i, j+1)
	res += maxAreaOfIslandHelper(grid, height, weight, i, j-1)
	return res
}
