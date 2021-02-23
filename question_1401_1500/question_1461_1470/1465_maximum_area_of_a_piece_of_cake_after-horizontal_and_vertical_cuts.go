package question_1461_1470

// 1465. 切割后面积最大的蛋糕
// https://leetcode-cn.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/
// Topics: 数组

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append(horizontalCuts, 0, h)
	verticalCuts = append(verticalCuts, 0, w)
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	a, b := 0, 0
	for i := 1; i < len(horizontalCuts); i++ {
		if horizontalCuts[i]-horizontalCuts[i-1] > b-a {
			a, b = horizontalCuts[i-1], horizontalCuts[i]
		}
	}

	c, d := 0, 0
	for i := 1; i < len(verticalCuts); i++ {
		if verticalCuts[i]-verticalCuts[i-1] > d-c {
			c, d = verticalCuts[i-1], verticalCuts[i]
		}
	}

	return (b - a) * (d - c) % 1000000007
}
