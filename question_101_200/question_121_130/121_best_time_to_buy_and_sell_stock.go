package question_121_130

// 121. 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

func maxProfit(prices []int) int {
	max := 0
	for i, p := range prices {
		for _, t := range prices[i+1:] {
			s := t - p
			if s > max {
				max = s
			}
		}
	}
	return max
}
