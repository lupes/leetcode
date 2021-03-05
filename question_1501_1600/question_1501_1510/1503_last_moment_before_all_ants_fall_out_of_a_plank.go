package question_1501_1510

// 1503. 所有蚂蚁掉下来前的最后一刻
// https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/
// Topics: 脑筋急转弯 数组

func getLastMoment(n int, left []int, right []int) int {
	max := 0
	for _, t := range left {
		if t > max {
			max = t
		}
	}

	for _, t := range right {
		if n-t > max {
			max = n - t
		}
	}
	return max
}
