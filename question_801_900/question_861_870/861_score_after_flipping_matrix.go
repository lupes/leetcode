package question_861_870

// 861. 翻转矩阵后的得分
// https://leetcode-cn.com/problems/score-after-flipping-matrix
// Topics: 贪心算法

func matrixScore(A [][]int) int {
	rl, cl, res, t := len(A), len(A[0]), 0, 0
	for i, row := range A {
		if row[0] == 0 {
			for j := range row {
				A[i][j] = 1 - A[i][j]
			}
		}
	}

	for i := 1; i < cl; i++ {
		var one = 0
		for j, _ := range A {
			if A[j][i] == 1 {
				one++
			}
		}
		if one*2 < rl {
			for j, _ := range A {
				A[j][i] = 1 - A[j][i]
			}
		}
	}

	for _, row := range A {
		t = 0
		for i, c := range row {
			t += c << (cl - i - 1)
		}
		res += t
	}
	return res
}
