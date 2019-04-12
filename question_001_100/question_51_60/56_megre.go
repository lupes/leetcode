package question_51_60

import "sort"

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	size := len(intervals)
	if size < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start > intervals[j].Start || (intervals[i].Start == intervals[j].Start && intervals[i].End > intervals[j].End) {
			return false
		}
		return true
	})
	var res = []Interval{intervals[0]}
	var l = 0
	for _, inter := range intervals {
		if inter.Start <= res[l].End {
			if inter.End > res[l].End {
				res[l].End = inter.End
			}
		} else {
			res = append(res, inter)
			l++
		}
	}
	return res
}
