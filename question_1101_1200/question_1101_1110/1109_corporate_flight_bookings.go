package question_1101_1110

// 1109. 航班预订统计
// https://leetcode-cn.com/problems/corporate-flight-bookings
// Topics: 数组 数学

func corpFlightBookings(bookings [][]int, n int) []int {
	var res = make([]int, n)
	for _, booking := range bookings {
		if booking[0] > n {
			continue
		}
		for i := booking[0] - 1; i < booking[1]; i++ {
			if i < n {
				res[i] += booking[2]
			}
		}
	}
	return res
}
