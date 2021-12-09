package question_791_800

// 794. 有效的井字游戏
// https://leetcode-cn.com/problems/valid-tic-tac-toe-state
// Topics: 递归 数学

func validTicTacToe(board []string) bool {
	o, os, x, xs := 0, 0, 0, 0
	for i, row := range board {

		// 同行
		switch row {
		case "OOO":
			os++
		case "XXX":
			xs++
		}

		// 同列
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			switch board[0][i] {
			case 'O':
				os++
			case 'X':
				xs++
			}
		}

		for _, c := range row {
			switch c {
			case 'O':
				o++
			case 'X':
				x++
			}
		}
	}

	// 对角线\
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		switch board[1][1] {
		case 'O':
			os++
		case 'X':
			xs++
		}
	}

	// 对角线/
	if board[2][0] == board[1][1] && board[1][1] == board[0][2] {
		switch board[1][1] {
		case 'O':
			os++
		case 'X':
			xs++
		}
	}

	if !(o == x || x == o+1) || (xs > 0 && o == x) || (os > 0 && x == o+1) {
		return false
	}

	if os > 0 && xs > 0 {
		return false
	}

	return true
}
