package question_321_330

// 322. 零钱兑换
// https://leetcode-cn.com/problems/coin-change
// Topics: 动态规划

func coinChange(coins []int, amount int) int {
	var res = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := -1
		for _, coin := range coins {
			t := i - coin
			if t >= 0 && res[t] > -1 && (min == -1 || res[t] < min) {
				min = res[t] + 1
			}
		}
		res[i] = min
	}
	return res[amount]
}
