package question_451_460

import (
	"sort"
)

// 452. 用最少数量的箭引爆气球
// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons
// Topics: 贪心算法

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] > points[j][1])
	})
	var res = 0
	for len(points) > 0 {
		a, b := points[0][0], points[0][1]
		for i, p := range points[1:] {
			if a <= p[0] && p[0] <= b {
				if a < p[0] {
					a = p[0]
				}
				if p[1] < b {
					b = p[1]
				}
			} else {
				points = points[i+1:]
				goto Next
			}
		}
		points = points[:0]
	Next:
		res += 1
	}
	return res
}
