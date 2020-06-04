package question_1181_1190

// 1184. 公交站间的距离
// https://leetcode-cn.com/problems/distance-between-bus-stops
// Topics: 数组

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if destination < start {
		start, destination = destination, start
	}
	var sum, tmp = 0, 0
	for i, n := range distance {
		sum += n
		if i >= start && i < destination {
			tmp += n
		}
	}
	if tmp > sum-tmp {
		return sum - tmp
	}
	return tmp
}
