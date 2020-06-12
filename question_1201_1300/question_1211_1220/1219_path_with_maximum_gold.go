package question_1211_1220

// 1219. 黄金矿工
// https://leetcode-cn.com/problems/path-with-maximum-gold
// Topics: 回溯算法

func getMaximumGold(grid [][]int) int {
	var max, res, tmp = 0, 0, 0
	for _, row := range grid {
		for _, cell := range row {
			max += cell
		}
	}
	for i, row := range grid {
		for j, cell := range row {
			if cell != 0 {
				tmp = getMaximumGoldHelp(grid, i, j, 0, max)
				if tmp == max {
					return max
				} else if tmp > res {
					res = tmp
				}
			}
		}
	}
	return res
}

func getMaximumGoldHelp(grid [][]int, i, j, sum, max int) int {
	var res, tmp, now, cell = sum + grid[i][j], 0, sum + grid[i][j], grid[i][j]

	grid[i][j] = 0
	if i > 0 && grid[i-1][j] > 0 {
		tmp = getMaximumGoldHelp(grid, i-1, j, now, max)
		if tmp == max {
			return max
		} else if tmp > res {
			res = tmp
		}
	}

	if i < len(grid)-1 && grid[i+1][j] > 0 {
		tmp = getMaximumGoldHelp(grid, i+1, j, now, max)
		if tmp == max {
			return max
		} else if tmp > res {
			res = tmp
		}
	}

	if j > 0 && grid[i][j-1] > 0 {
		tmp = getMaximumGoldHelp(grid, i, j-1, now, max)
		if tmp == max {
			return max
		} else if tmp > res {
			res = tmp
		}
	}

	if j < len(grid[0])-1 && grid[i][j+1] > 0 {
		tmp = getMaximumGoldHelp(grid, i, j+1, now, max)
		if tmp == max {
			return max
		} else if tmp > res {
			res = tmp
		}
	}
	grid[i][j] = cell
	return res
}
