package question_1041_1050

// 1049. 最后一块石头的重量 II
// https://leetcode-cn.com/problems/last-stone-weight-ii
// Topics: 动态规划

func lastStoneWeightII(stones []int) int {
	sum, tmp := 0, 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	var dp = make([]int, sum+1)
	for i := range stones {
		for j := target; j >= stones[i]; j-- {
			tmp = dp[j-stones[i]] + stones[i]
			if dp[j] < tmp {
				dp[j] = tmp
			}
		}
	}
	return sum - dp[target] - dp[target]
}
