package question_61_70

func uniquePaths(m int, n int) int {
	if m == 1 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}
