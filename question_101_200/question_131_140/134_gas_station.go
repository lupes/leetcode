package question_131_140

// 134. 加油站
// https://leetcode-cn.com/problems/gas-station
// Topics: 贪心算法

func canCompleteCircuit(gas []int, cost []int) int {
	res, run, start := 0, 0, 0
	for i, n := range gas {
		run += n - cost[i]
		res += n - cost[i]
		if run < 0 {
			start = i + 1
			run = 0
		}
	}
	if res < 0 {
		return -1
	}
	return start
}
