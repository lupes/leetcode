package question_1591_1600

// 1599. 经营摩天轮的最大利润
// https://leetcode-cn.com/problems/maximum-profit-of-operating-a-centennial-wheel/
// Topics: 贪心算法

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	var tmp, max, num, now, times = 0, 0, 0, 0, -1
	for i := 1; num > 0 || len(customers) > 0; i++ {
		if len(customers) > 0 {
			num, customers = num+customers[0], customers[1:]
		}
		if num > 4 {
			now, num = 4, num-4
		} else {
			now, num = num, 0
		}
		tmp = tmp + now*boardingCost - runningCost
		if tmp > max {
			max, times = tmp, i
		}
	}
	return times
}
