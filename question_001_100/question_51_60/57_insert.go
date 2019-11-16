package question_51_60

// 57. 插入区间
// https://leetcode-cn.com/problems/insert-interval

func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	if size == 0 {
		return [][]int{newInterval}
	}
	if len(newInterval) == 0 {
		return intervals
	}
	var res [][]int
	for i, inter := range intervals {
		if inter[0] > newInterval[1] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		}

		if inter[0] <= newInterval[0] && newInterval[1] <= inter[1] {
			res = append(res, intervals[i:]...)
			return res
		}

		if newInterval[0] <= inter[0] && newInterval[1] <= inter[1] {
			res = append(res, []int{newInterval[0], inter[1]})
			res = append(res, intervals[i+1:]...)
			return res
		}

		if newInterval[0] > inter[1] {
			res = append(res, inter)
			goto OVER
		}

		if inter[0] <= newInterval[0] && inter[1] < newInterval[1] {
			newInterval[0] = inter[0]
			goto OVER
		}
	OVER:
		if size-1 == i {
			res = append(res, newInterval)
			return res
		}
	}
	return res
}
