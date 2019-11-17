package qustion_511_520

// 518. 零钱兑换 II
// https://leetcode-cn.com/problems/coin-change-2/

func change(amount int, coins []int) int {
	var res = make([]int, amount+1)
	res[0] = 1
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if i-coin >= 0 {
				res[i] += res[i-coin]
			}
		}
	}
	return res[amount]
}
