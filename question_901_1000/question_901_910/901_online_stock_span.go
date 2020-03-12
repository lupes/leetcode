package question_901_910

// 901. 股票价格跨度
// https://leetcode-cn.com/problems/online-stock-span
// Topics: 栈

type StockSpanner struct {
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	return 0
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
