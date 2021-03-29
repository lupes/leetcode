package question_1571_1580

// 1578. 避免重复字母的最小删除成本
// https://leetcode-cn.com/problems/minimum-deletion-cost-to-avoid-repeating-letters/
// Topics: 贪心算法

func minCost(s string, cost []int) int {
	var res, last = 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[last] {
			if cost[i] > cost[last] {
				res += cost[last]
				last = i
			} else {
				res += cost[i]
			}
		} else {
			last = i
		}
	}
	return res
}
