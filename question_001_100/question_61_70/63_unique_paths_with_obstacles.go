package question_61_70

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	r, rl, c, cl := 0, len(obstacleGrid)-1, 0, len(obstacleGrid[0])-1
	if obstacleGrid[0][0] == 1 || obstacleGrid[rl][cl] == 1 {
		return 0
	}
	return uniquePathsWithObstaclesDfs(r, rl, c, cl, obstacleGrid)
}

func uniquePathsWithObstaclesDfs(r, rl, c, cl int, obstacleGrid [][]int) int {
	if r == rl && c == cl {
		return 1
	}
	res1, res2 := 0, 0
	if r < rl && obstacleGrid[r+1][c] != 1 {
		res1 = uniquePathsWithObstaclesDfs(r+1, rl, c, cl, obstacleGrid)
	}
	if c < cl && obstacleGrid[r][c+1] != 1 {
		res2 = uniquePathsWithObstaclesDfs(r, rl, c+1, cl, obstacleGrid)
	}
	return res1 + res2
}
