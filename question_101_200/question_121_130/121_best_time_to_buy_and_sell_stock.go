package question_121_130

// 121. 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
// Topics: 数组 动态规划

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy, sell := prices[0], 0
	for _, price := range prices[1:] {
		if price-buy > sell {
			sell = price - buy
		}
		if price < buy {
			buy = price
		}
	}
	return sell
}
