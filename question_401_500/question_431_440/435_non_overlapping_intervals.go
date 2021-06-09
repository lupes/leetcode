package question_431_440

import (
	"sort"
)

// 435. 无重叠区间
// https://leetcode-cn.com/problems/non-overlapping-intervals
// Topics: 贪心算法

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	now, res := intervals[0], 0
	for i := 1; i < len(intervals); i++ {
		if now[1] <= intervals[i][0] {
			now = intervals[i]
			continue
		}
		res += 1
		if now[1] > intervals[i][1] {
			now = intervals[i]
		}
	}
	return res
}
