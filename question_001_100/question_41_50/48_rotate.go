package question_41_50

func rotate(matrix [][]int) {
	size := len(matrix[0])
	for i := 0; i < size; i++ {
		l := size - i*2 - 1
		for j := 0; j < l; j++ {
			matrix[i][i+j], matrix[i+j][i+l], matrix[i+l][i+l-j], matrix[i+l-j][i] = matrix[i+l-j][i], matrix[i][i+j], matrix[i+j][i+l], matrix[i+l][i+l-j]
		}
	}
}
