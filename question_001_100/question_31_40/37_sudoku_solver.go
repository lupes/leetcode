package question_31_40

import "fmt"

// 37. 解数独
// https://leetcode-cn.com/problems/sudoku-solver
// Topics: 哈希表 回溯算法

func (b Board) String() string {
	res := ""
	for _, row := range b {
		line := ""
		for _, c := range row {
			line += fmt.Sprintf(", '%s'", string(c))
		}
		res += fmt.Sprintf("\t{%s},\n", line[2:])
	}
	return fmt.Sprintf("{\n%s}\n", res)
}

type Board [9][9]byte

func solveSudoku(board [][]byte) {
	var tmp Board
	for i, row := range board {
		for j, num := range row {
			tmp[i][j] = num
		}
	}
	res, _ := dfs(tmp, 0, 0)
	for i, row := range res {
		for j, num := range row {
			board[i][j] = num
		}
	}
}

func dfs(board Board, i, j int) (Board, bool) {
	jn := (j + 1) % 9
	in := i + (j+1)/9
	if i == 9 {
		return board, true
	}
	if board[i][j] != '.' {
		return dfs(board, in, jn)
	}

Next:
	for t := 1; t < 10; t++ {
		c := byte(t) + '0'
		for k := 0; k < 9; k++ {
			if board[i][k] == c {
				continue Next
			}
			if board[k][j] == c {
				continue Next
			}
			if board[i/3*3+k/3][j/3*3+k%3] == c {
				continue Next
			}
		}
		board[i][j] = c
		tmp, err := dfs(board, in, jn)
		if err {
			return tmp, true
		}
	}
	return board, false
}
