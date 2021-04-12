package question_1611_1620

import (
	"math"
)

// 1620. 网络信号最好的坐标
// https://leetcode-cn.com/problems/coordinate-with-maximum-network-quality/
// Topics: 贪心算法

func bestCoordinate(towers [][]int, radius int) []int {
	maxX, minX, maxY, minY := 0, 0, 0, 0
	for _, t := range towers {
		if t[0]+radius > maxX {
			maxX = t[0] + radius
		}
		if t[0]-radius >= 0 && t[0]-radius < minX {
			minX = t[0] - radius
		}
		if t[1]+radius > maxY {
			maxY = t[1] + radius
		}
		if t[1]-radius >= 0 && t[1]-radius < minY {
			minY = t[1] - radius
		}
	}

	var max = 0
	var res = []int{0, 0}
	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			var tmp = 0
			for _, t := range towers {
				a := (t[0]-i)*(t[0]-i) + (t[1]-j)*(t[1]-j)
				if a <= radius*radius {
					tmp += int(float64(t[2]) / (1.0 + math.Sqrt(float64(a))))
				}
			}
			//fmt.Printf("i:%d j:%d q:%d\n", i, j, tmp)
			if tmp > max {
				res, max = []int{i, j}, tmp
			}
		}
	}

	return res
}
