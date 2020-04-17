package question_981_990

// 983. 最低票价
// https://leetcode-cn.com/problems/minimum-cost-for-tickets
// Topics: 动态规划

func mincostTickets(days []int, costs []int) int {
	last := days[len(days)-1]
	var dp = make([]int, last+1)
	var flag = make(map[int]struct{})
	for _, n := range days {
		flag[n] = struct{}{}
	}
	for i := 1; i <= last; i++ {
		if _, ok := flag[i]; ok {
			dp[i] = min(
				costs[0]+dp[i-1],
				costs[1]+dp[max(0, i-7)],
				costs[2]+dp[max(0, i-30)])
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[last]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(arr ...int) int {
	min := arr[0]
	for _, a := range arr {
		if min > a {
			min = a
		}
	}
	return min
}
