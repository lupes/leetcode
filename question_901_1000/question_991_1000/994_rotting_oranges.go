package question_991_1000

// 994. 腐烂的橘子
// https://leetcode-cn.com/problems/rotting-oranges
// Topics: 广度优先搜索

func orangesRotting(grid [][]int) int {
	for i := 3; ; i++ {
		var rotting, norm = 0, 0
		for j := 0; j < len(grid); j++ {
			for k := 0; k < len(grid[j]); k++ {
				if grid[j][k] == 1 {
					norm++
				} else if grid[j][k] != 1 && grid[j][k] != 0 && grid[j][k] < i {
					if j > 0 && grid[j-1][k] == 1 {
						grid[j-1][k] = i
						rotting++
					}
					if j+1 < len(grid) && grid[j+1][k] == 1 {
						grid[j+1][k] = i
						rotting++
					}
					if k > 0 && grid[j][k-1] == 1 {
						grid[j][k-1] = i
						rotting++
					}
					if k+1 < len(grid[j]) && grid[j][k+1] == 1 {
						grid[j][k+1] = i
						rotting++
					}
				}
			}
		}
		if rotting == 0 && norm != 0 {
			return -1
		} else if rotting == 0 && norm == 0 {
			return i - 3
		}
	}
}
