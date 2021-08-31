package question_1101_1110

// 1109. 航班预订统计
// https://leetcode-cn.com/problems/corporate-flight-bookings
// Topics: 数组 数学 前缀和

func corpFlightBookings(bookings [][]int, n int) []int {
	var res = make([]int, n+2)
	for _, booking := range bookings {
		res[booking[0]] += booking[2]
		res[booking[1]+1] -= booking[2]
	}
	for i := 1; i < len(res); i++ {
		res[i] += res[i-1]
	}
	return res[1 : len(res)-1]
}
