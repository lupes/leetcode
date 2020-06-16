package question_1231_1240

// 1232. 缀点成线
// https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
// Topics: 几何 数组 数学

func checkStraightLine(coordinates [][]int) bool {
	p1, p2 := coordinates[0], coordinates[1]
	for _, p := range coordinates[2:] {
		if (p[0]-p1[0])*(p2[1]-p1[1])+p1[1]*(p2[0]-p1[0]) != p[1]*(p2[0]-p1[0]) {
			return false
		}
	}
	return true
}
