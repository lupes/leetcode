package question_111_120

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		var t = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				t[j] = 1
			} else {
				t[j] = res[i-1][j-1] + res[i-1][j]
			}
		}
		res[i] = t
	}
	return res
}
