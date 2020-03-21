package question_931_940

// 931. 下降路径最小和
// https://leetcode-cn.com/problems/minimum-falling-path-sum
// Topics: 动态规划

func minFallingPathSum(A [][]int) int {
	var res = make([][]int, len(A))
	rowLen := len(A[0])
	res[0] = make([]int, rowLen)
	for j, n := range A[0] {
		res[0][j] = n
	}
	for i := 1; i < len(A); i++ {
		res[i] = make([]int, rowLen)
		for j, n := range A[i] {
			min := res[i-1][j]
			if j > 0 && res[i-1][j-1] < min {
				min = res[i-1][j-1]
			}
			if j < rowLen-1 && res[i-1][j+1] < min {
				min = res[i-1][j+1]
			}
			res[i][j] = n + min
		}
	}
	min := res[rowLen-1][0]
	for _, n := range res[rowLen-1] {
		if n < min {
			min = n
		}
	}
	return min
}
