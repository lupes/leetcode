package question_901_910

// 901. 股票价格跨度
// https://leetcode-cn.com/problems/online-stock-span
// Topics: 栈

type Price struct {
	price int
	num   int
}

type StockSpanner struct {
	Prices []Price
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	var res = 1
	for i := len(this.Prices) - 1; i >= 0; {
		if this.Prices[i].price <= price {
			res += this.Prices[i].num
			i -= this.Prices[i].num
		} else {
			break
		}
	}
	this.Prices = append(this.Prices, Price{price: price, num: res})
	return res
}
