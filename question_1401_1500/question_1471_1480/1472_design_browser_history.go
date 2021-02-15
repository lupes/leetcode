package question_1471_1480

// 1472. 设计浏览器历史记录
// https://leetcode-cn.com/problems/design-browser-history/
// Topics: 设计

type BrowserHistory struct {
	now     string
	back    []string
	forward []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		now:     homepage,
		back:    nil,
		forward: nil,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.back = append(this.back, this.now)
	this.now = url
	this.forward = nil
}

func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps && 0 < len(this.back); i++ {
		this.forward = append(this.forward, this.now)
		this.now, this.back = this.back[len(this.back)-1], this.back[:len(this.back)-1]
	}
	return this.now
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps && 0 < len(this.forward); i++ {
		this.back = append(this.back, this.now)
		this.now, this.forward = this.forward[len(this.forward)-1], this.forward[:len(this.forward)-1]
	}
	return this.now
}
