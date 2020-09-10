package question_1281_1290

import (
	"sort"
)

// 1288. 删除被覆盖区间
// https://leetcode-cn.com/problems/remove-covered-intervals/
// Topics: Line Sweep

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var res = 1
	for now, i := 0, 1; i < len(intervals); i++ {
		if intervals[now][1] < intervals[i][1] {
			now, res = i, res+1
		}
	}
	return res
}
