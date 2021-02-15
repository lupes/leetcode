package question_1471_1480

// 1472. 设计浏览器历史记录
// https://leetcode-cn.com/problems/design-browser-history/
// Topics: 设计

type BrowserHistory struct {
	now     int
	history []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		now:     0,
		history: []string{homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.now++
	this.history = this.history[:this.now]
	this.history = append(this.history, url)
}

func (this *BrowserHistory) Back(steps int) string {
	this.now -= steps
	if this.now < 0 {
		this.now = 0
	}
	return this.history[this.now]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.now += steps
	if this.now >= len(this.history) {
		this.now = len(this.history) - 1
	}
	return this.history[this.now]
}
