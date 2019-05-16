package question_61_70

func minPathSum(grid [][]int) int {
	r, rl, c, cl, min := 1, len(grid), 1, len(grid[0]), 0
	for r = 1; r < rl; r++ {
		grid[r][0] += grid[r-1][0]
	}
	for c = 1; c < cl; c++ {
		grid[0][c] += grid[0][c-1]
	}
	for r = 1; r < rl; r++ {
		for c = 1; c < cl; c++ {
			min = grid[r][c-1]
			if min > grid[r-1][c] {
				min = grid[r-1][c]
			}
			grid[r][c] += min
		}
	}
	return grid[rl-1][cl-1]
}
