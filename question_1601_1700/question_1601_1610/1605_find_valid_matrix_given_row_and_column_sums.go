package question_1601_1610

// 1605. 给定行和列的和求可行矩阵
// https://leetcode-cn.com/problems/find-valid-matrix-given-row-and-column-sums/
// Topics: 字符串 OrderedMap

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	var res = make([][]int, len(rowSum))
	for i, row := range rowSum {
		res[i] = make([]int, len(colSum))
		for j, col := range colSum {
			t := row
			if col < t {
				t = col
			}
			res[i][j] = t
			row, colSum[j] = row-t, colSum[j]-t
		}
	}
	return res
}
