package question_841_850

// 849. 到最近的人的最大距离
// https://leetcode-cn.com/problems/maximize-distance-to-closest-person
// Topics: 数组

func maxDistToClosest(seats []int) int {
	var tmp []int
	for i, f := range seats {
		if f == 1 {
			tmp = append(tmp, i)
		}
	}
	var max = tmp[0]
	for i, n := range tmp {
		if len(tmp)-1 == i {
			if len(seats)-n-1 > max {
				max = len(seats) - n - 1
			}
		} else if (tmp[i+1]-n)/2 > max {
			max = (tmp[i+1] - n) / 2
		}
	}
	return max
}
