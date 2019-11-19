package question_121_130

// 122. 买卖股票的最佳时机 II
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
// Topics: 贪心算法 数组

func maxProfit2(prices []int) int {
	max, tmp, now, last := 0, 0, -1, -1
	for _, p := range prices {
		if p > last && last != -1 {
			tmp = p - now
			last = p
		} else {
			max += tmp
			tmp = 0
			now = p
			last = p
		}
	}
	return max + tmp
}
