package question_61_70

func minPathSum(grid [][]int) int {
	r, rl, c, cl := 0, len(grid)-1, 0, len(grid[0])-1
	var m = make(map[[2]int]int)
	return minPathSumDfs(r, rl, c, cl, grid, m)
}

func minPathSumDfs(r, rl, c, cl int, grid [][]int, m map[[2]int]int) int {
	res := -1
	res1, res2, ok := 0, 0, false
	if r < rl {
		res1, ok = m[[2]int{r + 1, c}]
		if !ok {
			res1 = minPathSumDfs(r+1, rl, c, cl, grid, m)
			m[[2]int{r + 1, c}] = res1
		}
		res = res1
	}
	if c < cl {
		res2, ok = m[[2]int{r, c + 1}]
		if !ok {
			res2 = minPathSumDfs(r, rl, c+1, cl, grid, m)
			m[[2]int{r, c + 1}] = res2
		}
		if res == -1 || res > res2 {
			res = res2
		}
	}
	if res == -1 {
		res = 0
	}
	return grid[r][c] + res
}
