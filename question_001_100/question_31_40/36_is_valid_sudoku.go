package question_31_40

func isValidSudoku(board [][]byte) bool {
	var flagsRow [9][9]bool
	var flagsCol [9][9]bool
	var flagsSqu [9][9]bool
	for i, row := range board {
		for j, num := range row {
			if num == '.' {
				continue
			}
			num = num - '1'
			k := (i/3)*3 + (j / 3)
			if flagsRow[i][num] || flagsCol[j][num] || flagsSqu[k][num] {
				return false
			}
			flagsRow[i][num] = true
			flagsCol[j][num] = true
			flagsSqu[k][num] = true
		}
	}
	return true
}
