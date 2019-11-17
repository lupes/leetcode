package question_281_290

// 289. 生命游戏
// https://leetcode-cn.com/problems/game-of-life/
// Topics: 数组

func gameOfLife(board [][]int) {

	ly := len(board)
	var flags = make([][]bool, ly)
	for y, ns := range board {
		lx := len(ns)
		var flag = make([]bool, lx)
		flags[y] = flag
		for x, n := range ns {
			t := 0
			if y > 0 && x > 0 && board[y-1][x-1] == 1 {
				t += 1
			}
			if y > 0 && board[y-1][x] == 1 {
				t += 1
			}
			if y > 0 && x+1 < lx && board[y-1][x+1] == 1 {
				t += 1
			}

			if x > 0 && board[y][x-1] == 1 {
				t += 1
			}
			if x+1 < lx && board[y][x+1] == 1 {
				t += 1
			}

			if y+1 < ly && x > 0 && board[y+1][x-1] == 1 {
				t += 1
			}
			if y+1 < ly && board[y+1][x] == 1 {
				t += 1
			}
			if y+1 < ly && x+1 < lx && board[y+1][x+1] == 1 {
				t += 1
			}

			if n == 1 && (t < 2 || t > 3) {
				flag[x] = true
			} else if n == 0 && t == 3 {
				flag[x] = true
			}
		}
	}

	for y, flag := range flags {
		for x, f := range flag {
			if f {
				if board[y][x] == 1 {
					board[y][x] = 0
				} else {
					board[y][x] = 1
				}
			}
		}
	}
}
