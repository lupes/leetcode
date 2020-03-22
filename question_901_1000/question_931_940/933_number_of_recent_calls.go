package question_931_940

// 933. 最近的请求次数
// https://leetcode-cn.com/problems/number-of-recent-calls
// Topics: 队列

type RecentCounter struct {
	data []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.data = append(this.data, t)
	var left, right = 0, len(this.data)
	for right > left {
		center := (left + right) / 2
		if t-this.data[center] > 3000 {
			left = center + 1
		} else {
			right = center
		}
	}
	return len(this.data[left:])
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
