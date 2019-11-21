package question_461_470

// 463. 岛屿的周长
// https://leetcode-cn.com/problems/island-perimeter
// Topics: 哈希表

func islandPerimeter(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	var res, h, w = 0, len(grid), len(grid[0])
	fun := helper(grid, h, w)
	for i, g := range grid {
		for j, n := range g {
			if n == 0 {
				continue
			}
			res += fun(i-1, j) + fun(i+1, j) + fun(i, j-1) + fun(i, j+1)
		}
	}
	return res
}

func helper(grid [][]int, height, wight int) func(row, col int) int {
	return func(row, col int) (i int) {
		if row < 0 || row >= height {
			return 1
		}
		if col < 0 || col >= wight {
			return 1
		}
		if grid[row][col] == 0 {
			return 1
		}
		return 0
	}
}
