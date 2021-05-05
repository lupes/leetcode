package question_1671_1680

// 1672. 最富有客户的资产总量
// https://leetcode-cn.com/problems/richest-customer-wealth/
// Topics: 数组

func maximumWealth(accounts [][]int) int {
	var max, tmp = 0, 0
	for _, customer := range accounts {
		tmp = 0
		for _, account := range customer {
			tmp += account
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
