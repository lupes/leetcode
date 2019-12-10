package question_121_130

// 122. 买卖股票的最佳时机 II
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
// Topics: 贪心算法 数组

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy, sell := prices[0], 0
	for _, price := range prices[1:] {
		if price-sell < buy {
			buy = price - sell
		}
		if price-buy > sell {
			sell = price - buy
		}
	}
	return sell
}
