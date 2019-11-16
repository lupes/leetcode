package question_51_60

// 54. 螺旋矩阵
// https://leetcode-cn.com/problems/spiral-matrix

func spiralOrder(matrix [][]int) []int {
	size := len(matrix)
	if size == 0 {
		return nil
	}
	l, t, r, b := 0, 0, len(matrix[0])-1, size-1
	var res = make([]int, 0, (r+1)*(b+1))
	for l <= r || t <= b {
		for i := l; i <= r && b >= t; i++ {
			res = append(res, matrix[t][i])
		}
		for i := t + 1; i <= b && r >= l; i++ {
			res = append(res, matrix[i][r])
		}
		for i := r - 1; i >= l && b > t; i-- {
			res = append(res, matrix[b][i])
		}
		for i := b - 1; i >= t+1 && r > l; i-- {
			res = append(res, matrix[i][l])
		}
		l, t, r, b = l+1, t+1, r-1, b-1
	}
	return res
}
