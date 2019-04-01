package question_31_40

func isValidSudoku(board [][]byte) bool {
	var flagsLine [9][9]bool
	var flagsCol [9][9]bool
	var flags [9][9]bool
	for i, line := range board {
		for j, num := range line {
			if num == '.' {
				continue
			}
			num = num - '1'
			k := (i/3)*3 + (j / 3)
			if flagsLine[i][num] || flagsCol[j][num] || flags[k][num] {
				return false
			} else {
				flagsLine[i][num] = true
				flagsCol[j][num] = true
				flags[k][num] = true
			}
		}
	}
	return true
}
