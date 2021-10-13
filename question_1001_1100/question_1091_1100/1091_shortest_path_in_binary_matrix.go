package question_1091_1100

// 1091. 二进制矩阵中的最短路径
// https://leetcode-cn.com/problems/shortest-path-in-binary-matrix
// Topics: 广度优先搜索

var direction = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	if n == 1 {
		return n
	}
	grid[0][0] = 2
	var now, res = [][2]int{{0, 0}}, 1
	for len(now) > 0 {
		var next [][2]int
		for _, point := range now {
			for _, d := range direction {
				i, j := point[0]+d[0], point[1]+d[1]
				if i >= 0 && i < n && j >= 0 && j < n && grid[i][j] == 0 {
					if i == n-1 && j == n-1 {
						return res + 1
					}
					grid[i][j] = 2
					next = append(next, [2]int{i, j})
				}
			}
		}
		now = next
		res++
	}

	return -1
}
