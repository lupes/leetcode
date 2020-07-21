package question_521_530

// 529. 扫雷游戏
// https://leetcode-cn.com/problems/minesweeper
// Topics: 深度优先搜索 广度优先搜索

var boardFlag = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	updateBoardHelper(board, click, len(board), len(board[0]))
	return board
}

func updateBoardHelper(board [][]byte, click []int, cl, rl int) {
	var t, a, b = byte(0), click[0], click[1]
	for _, f := range boardFlag {
		x, y := click[0]+f[0], click[1]+f[1]
		if x >= 0 && y >= 0 && x < cl && y < rl {
			if board[x][y] == 'M' {
				t++
			}
		}
	}
	if t > 0 {
		board[a][b] = '0' + t
	} else if t == 0 {
		board[a][b] = 'B'
		for _, f := range boardFlag {
			x, y := click[0]+f[0], click[1]+f[1]
			if x >= 0 && y >= 0 && x < cl && y < rl {
				if board[x][y] == 'E' {
					updateBoardHelper(board, []int{x, y}, cl, rl)
				}
			}
		}
	}
}
