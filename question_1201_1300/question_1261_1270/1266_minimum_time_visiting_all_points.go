package question_1261_1270

// 1266. 访问所有点的最小时间
// https://leetcode-cn.com/problems/minimum-time-visiting-all-points/
// Topics: 几何 数组

func minTimeToVisitAllPoints(points [][]int) int {
	var start = points[0]
	var res = 0
	for i := 1; i < len(points); i++ {
		x := points[i][0] - start[0]
		y := points[i][1] - start[1]
		if x < 0 {
			x *= -1
		}
		if y < 0 {
			y *= -1
		}
		if x > y {
			res += x
		} else {
			res += y
		}
		start = points[i]
	}
	return res
}
