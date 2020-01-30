package question_741_750

// 746. 使用最小花费爬楼梯
// https://leetcode-cn.com/problems/min-cost-climbing-stairs
// Topics: 数组 动态规划

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	var res = make([]int, len(cost))
	res[0], res[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		res[i] = res[i-2] + cost[i]
		if res[i-1]+cost[i] < res[i] {
			res[i] = res[i-1] + cost[i]
		}
	}
	return res[len(cost)-1]
}
