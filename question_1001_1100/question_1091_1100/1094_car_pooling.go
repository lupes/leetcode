package question_1091_1100

// 1094. 拼车
// https://leetcode-cn.com/problems/car-pooling
// Topics: 贪心算法

func carPooling(trips [][]int, capacity int) bool {
	var flag = make(map[int]int)
	var min, max = 1000, 0
	for _, trip := range trips {
		if trip[1] < min {
			min = trip[1]
		}
		if trip[2] > max {
			max = trip[2]
		}
		flag[trip[1]] += trip[0]
		flag[trip[2]] -= trip[0]
	}
	for i := min; i < max; i++ {
		capacity -= flag[i]
		if capacity < 0 {
			return false
		}
	}
	return true
}
