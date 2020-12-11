package question_31_40

// 37. 解数独
// https://leetcode-cn.com/problems/sudoku-solver
// Topics: 哈希表 回溯算法

func solveSudoku(board [][]byte) {
	var flagsRow [9][9]bool
	var flagsCol [9][9]bool
	var flagsSqu [9][9]bool
	for i, row := range board {
		for j, num := range row {
			if num != '.' {
				flagsRow[i][num-'1'] = true
				flagsCol[j][num-'1'] = true
				flagsSqu[i/3*3+j/3][num-'1'] = true
			}
		}
	}
	solveSudokuHelper(board, flagsRow, flagsCol, flagsSqu, 0, 0)
}

func solveSudokuHelper(board [][]byte, flagsRow, flagsCol, flagsSqu [9][9]bool, i, j int) bool {
	if i == 8 && j == 9 {
		return true
	} else if j == 9 {
		j, i = 0, i+1
	}

	if board[i][j] != '.' {
		return solveSudokuHelper(board, flagsRow, flagsCol, flagsSqu, i, j+1)
	}

	for c := byte(0); c < 9; c++ {
		if flagsRow[i][c] || flagsCol[j][c] || flagsSqu[i/3*3+j/3][c] {
			continue
		}
		flagsRow[i][c], flagsCol[j][c], flagsSqu[i/3*3+j/3][c] = true, true, true
		board[i][j] = c + '1'
		if solveSudokuHelper(board, flagsRow, flagsCol, flagsSqu, i, j+1) {
			return true
		}
		flagsRow[i][c], flagsCol[j][c], flagsSqu[i/3*3+j/3][c] = false, false, false
		board[i][j] = '.'
	}

	return false
}
