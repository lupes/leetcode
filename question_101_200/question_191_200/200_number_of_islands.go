package question_191_200

// 200. 岛屿数量
// https://leetcode-cn.com/problems/number-of-islands
// Topics: 深度优先搜索 广度优先搜索 并查集

func numIslands(grid [][]byte) int {
	var res, rl, cl = 0, len(grid), len(grid[0])
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				sameLand(grid, rl, cl, i, j)
				res++
			}
		}
	}
	return res
}

func sameLand(grid [][]byte, rl, cl, i, j int) {
	grid[i][j] = 0
	if i > 0 && grid[i-1][j] == '1' {
		sameLand(grid, rl, cl, i-1, j)
	}
	if i < rl-1 && grid[i+1][j] == '1' {
		sameLand(grid, rl, cl, i+1, j)
	}

	if j > 0 && grid[i][j-1] == '1' {
		sameLand(grid, rl, cl, i, j-1)
	}
	if j < cl-1 && grid[i][j+1] == '1' {
		sameLand(grid, rl, cl, i, j+1)
	}
}
