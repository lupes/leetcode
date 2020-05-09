package question_1021_1030

import (
	"sort"
)

// 1030. 距离顺序排列矩阵单元格
// https://leetcode-cn.com/problems/matrix-cells-in-distance-order
// Topics: 排序

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var res = make([][]int, 0)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return helper(r0, c0, res[i][0], res[i][1]) < helper(r0, c0, res[j][0], res[j][1])
	})
	return res
}

func helper(x1, y1, x2, y2 int) int {
	x, y := x1-x2, y1-y2
	if x2 > x1 {
		x = x2 - x1
	}
	if y2 > y1 {
		y = y2 - y1
	}
	return x + y
}
