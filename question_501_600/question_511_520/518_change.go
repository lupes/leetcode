package qustion_511_520

// 518. 零钱兑换 II
// https://leetcode-cn.com/problems/coin-change-2/

func change(amount int, coins []int) int {
	if len(coins) == 0 && amount != 0 {
		return 0
	} else if amount == 0 {
		return 1
	}
	coin := coins[0]
	res := 0
	for i := 0; i <= amount/coin; i++ {
		res += change(amount-i*coin, coins[1:])
	}
	return res
}
