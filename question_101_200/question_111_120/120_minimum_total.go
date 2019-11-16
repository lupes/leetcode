package question_111_120

// 120. 三角形最小路径和
// https://leetcode-cn.com/problems/triangle

func minimumTotal(triangle [][]int) int {
	size := len(triangle)
	var res = triangle[0][0]
	for i := 1; i < size; i++ {
		res = 1<<63 - 1
		for j, n := range triangle[i] {
			l, r := 1<<63-1, 1<<63-1
			if j-1 >= 0 {
				l = triangle[i-1][j-1]
			}
			if j < len(triangle[i-1]) {
				r = triangle[i-1][j]
			}
			if r < l {
				l = r
			}
			triangle[i][j] = n + l
			if triangle[i][j] < res {
				res = triangle[i][j]
			}
		}
	}
	return res
}
