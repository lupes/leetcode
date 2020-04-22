package question_991_1000

// 994. 腐烂的橘子
// https://leetcode-cn.com/problems/rotting-oranges
// Topics: 广度优先搜索

func orangesRotting(grid [][]int) int {
	m := 0
	r := orangesRottingNum(grid)
	if r == 0 {
		return m
	}
	for {
		m++
		for i, line := range grid {
			for j, row := range line {
				if grid[i][j] == 0 {

				}
				if i > 0 && j > 0 && grid[i-1][j-1] == 1 {

				}
				if row == 1 {
					r++
				}
			}
		}
	}
}

func orangesRottingNum(grid [][]int) int {
	r := 0
	for _, line := range grid {
		for _, row := range line {
			if row == 1 {
				r++
			}
		}
	}
	return r
}
