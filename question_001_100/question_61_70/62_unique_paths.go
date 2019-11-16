package question_61_70

// 62. 不同路径
// https://leetcode-cn.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	if n > m {
		n, m = m, n
	}
	r1, r2 := 1, 1
	for i := m; i < m+n-1; i++ {
		r1 *= i
	}
	for i := 2; i < n; i++ {
		r2 *= i
	}
	return r1 / r2
}
