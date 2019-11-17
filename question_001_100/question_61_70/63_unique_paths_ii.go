package question_61_70

// 63. 不同路径 II
// https://leetcode-cn.com/problems/unique-paths-ii
// Topics: 数组 动态规划

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := make(map[[2]int]int)
	r, rl, c, cl := 0, len(obstacleGrid)-1, 0, len(obstacleGrid[0])-1
	if obstacleGrid[0][0] == 1 || obstacleGrid[rl][cl] == 1 {
		return 0
	}
	return uniquePathsWithObstaclesDfs(r, rl, c, cl, obstacleGrid, m)
}

func uniquePathsWithObstaclesDfs(r, rl, c, cl int, obstacleGrid [][]int, m map[[2]int]int) int {
	if r == rl && c == cl && obstacleGrid[rl][cl] != 1 {
		return 1
	}
	res1, res2, ok := 0, 0, false
	if r < rl && obstacleGrid[r+1][c] != 1 {
		res1, ok = m[[2]int{r + 1, c}]
		if !ok {
			res1 = uniquePathsWithObstaclesDfs(r+1, rl, c, cl, obstacleGrid, m)
			m[[2]int{r + 1, c}] = res1
		}
	}
	if c < cl && obstacleGrid[r][c+1] != 1 {
		res2, ok = m[[2]int{r, c + 1}]
		if !ok {
			res2 = uniquePathsWithObstaclesDfs(r, rl, c+1, cl, obstacleGrid, m)
			m[[2]int{r, c + 1}] = res2
		}
	}
	return res1 + res2
}
