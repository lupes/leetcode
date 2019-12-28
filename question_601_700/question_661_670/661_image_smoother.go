package question_661_670

// 661. 图片平滑器
// https://leetcode-cn.com/problems/image-smoother
// Topics: 数组

func imageSmoother(M [][]int) [][]int {
	rl := len(M)
	if rl < 1 {
		return nil
	}
	cl := len(M[0])
	var res = make([][]int, rl)
	for i, row := range M {
		res[i] = make([]int, cl)
		for j, n := range row {
			sum, num := n, 1
			if i-1 >= 0 && j-1 >= 0 {
				sum, num = sum+M[i-1][j-1], num+1
			}
			if i-1 >= 0 {
				sum, num = sum+M[i-1][j], num+1
			}
			if i-1 >= 0 && j+1 < cl {
				sum, num = sum+M[i-1][j+1], num+1
			}
			if j-1 >= 0 {
				sum, num = sum+M[i][j-1], num+1
			}
			if j+1 < cl {
				sum, num = sum+M[i][j+1], num+1
			}
			if i+1 < rl {
				sum, num = sum+M[i+1][j], num+1
			}
			if i+1 < rl && j-1 >= 0 {
				sum, num = sum+M[i+1][j-1], num+1
			}
			if i+1 < rl && j+1 < cl {
				sum, num = sum+M[i+1][j+1], num+1
			}
			res[i][j] = sum / num
		}
	}
	return res
}
