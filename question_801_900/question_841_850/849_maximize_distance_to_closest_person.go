package question_841_850

// 849. 到最近的人的最大距离
// https://leetcode-cn.com/problems/maximize-distance-to-closest-person
// Topics: 数组

func maxDistToClosest(seats []int) int {
	var last = -1
	var max = 0
	for i, f := range seats {
		if f == 1 {
			if last == -1 {
				max = i
			} else if (i-last)/2 > max {
				max = (i - last) / 2
			}
			last = i
		}
	}
	if len(seats)-last-1 > max {
		max = len(seats) - last - 1
	}
	return max
}
