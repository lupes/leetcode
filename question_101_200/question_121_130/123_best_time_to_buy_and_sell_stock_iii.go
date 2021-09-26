package question_121_130

import (
	"math"
)

// 123. 买卖股票的最佳时机 III
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
// Topics: 数组 动态规划

func maxProfit3(prices []int) int {
	var buy1, sell1, buy2, sell2 = math.MinInt32, 0, math.MinInt32, 0
	for _, price := range prices {
		if sell2 < buy2+price {
			sell2 = buy2 + price
		}
		if buy2 < sell1-price {
			buy2 = sell1 - price
		}

		if sell1 < buy1+price {
			sell1 = buy1 + price
		}
		if buy1 < -price {
			buy1 = -price
		}
	}
	return sell2
}
