package question_861_870

// 867. 转置矩阵
// https://leetcode-cn.com/problems/transpose-matrix
// Topics: 数组

func transpose(A [][]int) [][]int {
	var a = make([][]int, len(A[0]))
	for i := range a {
		a[i] = make([]int, len(A))
	}
	for i, row := range A {
		for j, num := range row {
			a[j][i] = num
		}
	}
	return a
}
