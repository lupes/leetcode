package question_1031_1040

// 1037. 有效的回旋镖
// https://leetcode-cn.com/problems/valid-boomerang
// Topics: 数学

func isBoomerang(points [][]int) bool {
	return (points[2][0]-points[1][0])*(points[1][1]-points[0][1]) != (points[2][1]-points[1][1])*(points[1][0]-points[0][0])
}
