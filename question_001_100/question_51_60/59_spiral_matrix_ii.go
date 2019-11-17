package question_51_60

// 59. 螺旋矩阵 II
// https://leetcode-cn.com/problems/spiral-matrix-ii

func generateMatrix(n int) [][]int {
	var res = make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}
	c, i := 0, 1
	for i <= n*n {
		for j := 0; j < n-2*c; j++ {
			res[c][j+c] = i
			i++
		}
		for j := 0; j < n-2*c-1; j++ {
			res[j+c+1][n-c-1] = i
			i++
		}
		for j := 0; j < n-2*c-1; j++ {
			res[n-c-1][n-c-j-2] = i
			i++
		}
		for j := 0; j < n-2*c-2; j++ {
			res[n-c-1-j-1][c] = i
			i++
		}
		c++
	}
	return res
}
