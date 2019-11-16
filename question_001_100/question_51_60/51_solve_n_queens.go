package question_51_60

// 51. N皇后
// https://leetcode-cn.com/problems/n-queens

func solveNQueens(n int) [][]string {
	flag := make([][]bool, n)
	for i := range flag {
		flag[i] = make([]bool, n)
	}
	return solveNQueensDfs(flag, 0, n)
}

func solveNQueensDfs(flag [][]bool, i, n int) (res [][]string) {
	for j := 0; j < n; j++ {
		flag[i][j] = true
		if valid(flag, i, j, n) {
			if i < n-1 {
				res = append(res, solveNQueensDfs(flag, i+1, n)...)
			} else {
				res = append(res, parse(flag))
			}
		}
		flag[i][j] = false
	}
	return
}

func parse(flag [][]bool) (res []string) {
	for _, arr := range flag {
		t := ""
		for _, a := range arr {
			if a {
				t += "Q"
			} else {
				t += "."
			}
		}
		res = append(res, t)
	}
	return res
}

func valid(flag [][]bool, i, j, n int) bool {
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
