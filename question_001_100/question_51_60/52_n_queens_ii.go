package question_51_60

// 52. N皇后 II
// https://leetcode-cn.com/problems/n-queens-ii

func totalNQueens(n int) int {
	flag := make([][]bool, n)
	for i := range flag {
		flag[i] = make([]bool, n)
	}
	return totalNQueensDfs(flag, 0, n)
}

func totalNQueensDfs(flag [][]bool, i, n int) (res int) {
	for j := 0; j < n; j++ {
		flag[i][j] = true
		if valid2(flag, i, j, n) {
			if i < n-1 {
				res += totalNQueensDfs(flag, i+1, n)
			} else {
				res += 1
			}
		}
		flag[i][j] = false
	}
	return
}

func valid2(flag [][]bool, i, j, n int) bool {
	for k := 0; k < i; k++ {
		if flag[k][j] {
			return false
		}
		if j-1-k >= 0 && flag[i-1-k][j-1-k] {
			return false
		}
		if j+1+k < n && flag[i-1-k][j+1+k] {
			return false
		}
	}
	return true
}
