package question_1251_1260

// 1253. 重构 2 行二进制矩阵
// https://leetcode-cn.com/problems/reconstruct-a-2-row-binary-matrix
// Topics: 贪心算法 数学

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var res = make([][]int, 2)
	res[0] = make([]int, len(colsum))
	res[1] = make([]int, len(colsum))
	a1, a2 := 0, 0
	for i, c := range colsum {
		if c == 0 {
			res[0][i], res[1][i] = 0, 0
		} else if c == 2 {
			res[0][i], res[1][i] = 1, 1
			upper, lower = upper-1, lower-1
		} else {
			a1, a2 = a1+1, a2+1
		}
	}
	if a1+a2 != (upper+lower)*2 || upper < 0 || lower < 0 {
		return nil
	}
	for i, c := range colsum {
		if c == 1 && upper > 0 {
			res[0][i], upper = 1, upper-1
		} else if c == 1 && lower > 0 {
			res[1][i], lower = 1, lower-1
		}
	}
	return res
}
