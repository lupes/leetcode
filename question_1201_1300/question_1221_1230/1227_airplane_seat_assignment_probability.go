package question_1221_1230

// 1227. 飞机座位分配概率
// https://leetcode-cn.com/problems/airplane-seat-assignment-probability
// Topics: 脑筋急转弯 数学 动态规划

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	} else {
		return 0.5
	}
}
