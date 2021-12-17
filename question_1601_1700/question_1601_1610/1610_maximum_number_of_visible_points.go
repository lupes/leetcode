package question_1601_1610

import (
	"math"
	"sort"
)

// 1610. 可见点的最大数目
// https://leetcode-cn.com/problems/maximum-number-of-visible-points/
// Topics: 几何 数组 数学 排序 滑动窗口

func visiblePoints(points [][]int, angle int, location []int) int {
	var same = 0
	var f []float64
	for _, point := range points {
		if point[0] == location[0] && point[1] == location[1] {
			same++
		} else {
			f = append(f, math.Atan2(float64(point[1]-location[1]), float64(point[0]-location[0])))
		}
	}
	sort.Float64s(f)

	n := len(f)
	for _, deg := range f {
		f = append(f, deg+2*math.Pi)
	}

	maxCnt := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range f[:n] {
		j := sort.Search(n*2, func(j int) bool { return f[j] > deg+degree })
		if j-i > maxCnt {
			maxCnt = j - i
		}
	}
	return same + maxCnt
}
