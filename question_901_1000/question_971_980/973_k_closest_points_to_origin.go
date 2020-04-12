package question_971_980

import (
	"sort"
)

// 973. 最接近原点的 K 个点
// https://leetcode-cn.com/problems/k-closest-points-to-origin
// Topics: 堆 排序 分治算法

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return (points[i][0]*points[i][0] + points[i][1]*points[i][1]) <
			(points[j][0]*points[j][0] + points[j][1]*points[j][1])
	})
	return points[:K]
}
