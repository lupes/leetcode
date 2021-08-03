package question_121_130

// 130. 被围绕的区域
// https://leetcode-cn.com/problems/surrounded-regions
// Topics: 深度优先搜索 广度优先搜索 并查集

var diff = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func solve(board [][]byte) {
	rl, cl := len(board), len(board[0])

	var tmp [][2]int
	for i := 0; i < rl; i++ {
		if board[i][0] == 'O' {
			tmp = append(tmp, [2]int{i, 0})
		}
		if board[i][cl-1] == 'O' {
			tmp = append(tmp, [2]int{i, cl - 1})
		}
	}
	for i := 0; i < cl; i++ {
		if board[0][i] == 'O' {
			tmp = append(tmp, [2]int{0, i})
		}
		if board[rl-1][i] == 'O' {
			tmp = append(tmp, [2]int{rl - 1, i})
		}
	}

	for _, key := range tmp {
		solveHelper(board, rl, cl, key[0], key[1])
	}

	for r, row := range board {
		for c, col := range row {
			if col == 'F' {
				board[r][c] = 'O'
			} else if col == 'O' {
				board[r][c] = 'X'
			}
		}
	}
}

func solveHelper(board [][]byte, rl, cl, r, c int) {
	if r < 0 || r >= rl || c < 0 || c >= cl {
		return
	}
	if board[r][c] == 'X' || board[r][c] == 'F' {
		return
	}
	board[r][c] = 'F'
	for _, d := range diff {
		solveHelper(board, rl, cl, r+d[0], c+d[1])
	}
}
