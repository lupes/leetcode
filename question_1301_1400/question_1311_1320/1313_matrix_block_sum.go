package question_1311_1320

// 1314. 矩阵区域和
// https://leetcode-cn.com/problems/matrix-block-sum/
// Topics: 动态规划

func matrixBlockSum(mat [][]int, K int) [][]int {
	var dp1 = make([][]int, len(mat)+K)
	var dp2 = make([][]int, len(mat)+K)
	for i := range dp1 {
		dp1[i] = make([]int, len(mat[0])+K)
		dp2[i] = make([]int, len(mat[0])+K)
	}

	for i := range mat {
		for j := range dp1[i] {
			if j == 0 {
				dp1[i][j] = mat[i][j]
			} else if j < len(mat[i]) {
				dp1[i][j] = mat[i][j] + dp1[i][j-1]
			} else {
				dp1[i][j] = dp1[i][j-1]
			}
			if j-2*K-1 >= 0 {
				dp1[i][j] -= mat[i][j-2*K-1]
			}
		}
	}

	for i, row := range dp1 {
		for j, n := range row {
			if i == 0 {
				dp2[i][j] = n
			} else {
				dp2[i][j] = n + dp2[i-1][j]
			}
			if i-2*K-1 >= 0 {
				dp2[i][j] -= dp1[i-2*K-1][j]
			}
		}
	}

	var res [][]int
	for i := K; i < len(dp2); i++ {
		res = append(res, dp2[i][K:])
	}

	return res
}
