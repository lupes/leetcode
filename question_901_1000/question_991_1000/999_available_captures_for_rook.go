package question_991_1000

// 999. 车的可用捕获量
// https://leetcode-cn.com/problems/available-captures-for-rook
// Topics: 数组

func numRookCaptures(board [][]byte) int {
	i, j := findRook(board)
	var res = 0
	for k := i - 1; k >= 0; k-- {
		if board[k][j] == 'p' {
			res++
			break
		} else if board[k][j] == 'B' {
			break
		}
	}
	for k := i + 1; k < 8; k++ {
		if board[k][j] == 'p' {
			res++
			break
		} else if board[k][j] == 'B' {
			break
		}
	}
	for k := j - 1; k >= 0; k-- {
		if board[i][k] == 'p' {
			res++
			break
		} else if board[i][k] == 'B' {
			break
		}
	}
	for k := j + 1; k < 8; k++ {
		if board[i][k] == 'p' {
			res++
			break
		} else if board[i][k] == 'B' {
			break
		}
	}
	return res
}

func findRook(board [][]byte) (int, int) {
	for i, line := range board {
		for j, v := range line {
			if v == 'R' {
				return i, j
			}
		}
	}
	return -1, -1
}
