package question_221_230

// 221. 最大正方形
// https://leetcode-cn.com/problems/maximal-square
// Topics: 动态规划

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix)+1)
	dp[0] = make([]int, len(matrix[0])+1)
	res := 0
	for i, row := range matrix {
		dp[i+1] = make([]int, len(row)+1)
		for j, col := range row {
			if col == '1' {
				t := dp[i][j]
				if dp[i][j+1] < t {
					t = dp[i][j+1]
				}
				if dp[i+1][j] < t {
					t = dp[i+1][j]
				}
				t += 1
				dp[i+1][j+1] = t
				if t*t > res {
					res = t * t
				}
			}
		}
	}
	return res
}
