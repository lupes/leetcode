package question_1011_1020

// 1020. 飞地的数量
// https://leetcode-cn.com/problems/number-of-enclaves
// Topics: 深度优先搜索

func numEnclaves(A [][]int) int {
	var res, f, rl, cl = 0, 2, len(A), len(A[0])
	for r, row := range A {
		for c, col := range row {
			if col == 1 {
				f++
				n, t := enclaves(A, rl, cl, r, c, f, 0)
				if !t {
					res += n
				}
			}
		}
	}
	return res
}

func enclaves(A [][]int, rl, cl, r, c, f, n int) (int, bool) {
	if r == rl || r < 0 || c == cl || c < 0 {
		return n, true
	}
	if A[r][c] != 1 {
		return n, false
	}

	n++
	A[r][c] = f
	n, t1 := enclaves(A, rl, cl, r-1, c, f, n)
	n, t2 := enclaves(A, rl, cl, r+1, c, f, n)
	n, t3 := enclaves(A, rl, cl, r, c-1, f, n)
	n, t4 := enclaves(A, rl, cl, r, c+1, f, n)
	return n, t1 || t2 || t3 || t4
}
