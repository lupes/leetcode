package question_901_910

// 901. 股票价格跨度
// https://leetcode-cn.com/problems/online-stock-span
// Topics: 栈

type StockSpanner struct {
	Prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	var res = 1
	for i := len(this.Prices) - 1; i >= 0; i-- {
		if this.Prices[i] <= price {
			res++
		} else {
			break
		}
	}
	this.Prices = append(this.Prices, price)
	return res
}
