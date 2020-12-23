package question_831_840

// 832. 翻转图像
// https://leetcode-cn.com/problems/flipping-an-image
// Topics: 数组

func flipAndInvertImage(A [][]int) [][]int {
	rl, cl := len(A), len(A[0])
	for i := 0; i < rl; i++ {
		for j := 0; j < cl; j++ {
			if j < cl/2 {
				A[i][j], A[i][cl-j-1] = A[i][cl-j-1], A[i][j]
			}
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}
	return A
}
